package cmd

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "bitbucket.org/envirovisionsolutions/showandtell/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"

	auth "bitbucket.org/envirovisionsolutions/showandtell/auth"
	database "bitbucket.org/envirovisionsolutions/showandtell/database"
	utils "bitbucket.org/envirovisionsolutions/showandtell/utils"
	web "bitbucket.org/envirovisionsolutions/showandtell/web"
)

type client struct{} // Add more data to this type if needed

var clients = make(map[*websocket.Conn]client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = client{}
			log.Println("connection registered")

		case message := <-broadcast:
			log.Println("message received:", message)

			// Send the message to all clients
			for connection := range clients {
				if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)

					unregister <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("connection unregistered")
		}
	}
}

func Serve() {
	engine := html.NewFileSystem(http.FS(utils.CmdContext.IndexFile), ".html")
	appName := utils.EnvGet("APP_NAME", "Show and Tell")
	isDevelopment := utils.EnvGet("APP_ENV", "production") == "development"
	app := fiber.New(fiber.Config{
		Prefork:           false, // Not compatible with websockets
		AppName:           appName,
		Views:             engine,
		EnablePrintRoutes: isDevelopment,
	})

	// Recover from panic middleware
	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC3339,
		Format:     "[${time}] (${ip}) Took: ${latency} ${status} - ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return isDevelopment
		},
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	go runHub()

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer func() {
			unregister <- c
			c.Close()
		}()

		// Register the client
		register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				// Broadcast the received message
				log.Printf("broadcasting message: %s\n", message)
				broadcast <- string(message)
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	cacheLengthMinutes, err := strconv.ParseInt(utils.EnvGet("APP_WEB_CACHE_LENGTH_MINUTES", "0"), 10, 32)
	if err != nil {
		cacheLengthMinutes = 0
	}
	if cacheLengthMinutes > 0 {
		log.Printf("INFO: Web route caching active. Timeout (minutes): %d\n", cacheLengthMinutes)
		app.Use(cache.New(cache.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.Query("refresh") == "true" || strings.Contains(c.Route().Path, "/ws")
			},
			Expiration:   time.Duration(cacheLengthMinutes) * time.Minute,
			CacheControl: true,
		}))
	}

	database.ConnectDb()
	keycloak := auth.NewKeycloak()
	web.SetupRoutes(app, keycloak)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(utils.CmdContext.StaticDir),
		PathPrefix: "web/static/src/dist",
		Browse:     true,
	}))

	app.Use(favicon.New(favicon.Config{
		URL:        "/favicon.ico",
		File:       "web/static/src/public/favicon.ico",
		FileSystem: http.FS(utils.CmdContext.FavIcon),
	}))

	port := utils.EnvGet("APP_PORT", "8021")
	log.Fatal(app.Listen(":" + port))
}
