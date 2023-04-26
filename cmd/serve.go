package cmd

import (
	"log"
	"net/http"
	"strconv"
	"time"

	_ "bitbucket.org/envirovisionsolutions/showandtell/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	auth "bitbucket.org/envirovisionsolutions/showandtell/auth"
	database "bitbucket.org/envirovisionsolutions/showandtell/database"
	utils "bitbucket.org/envirovisionsolutions/showandtell/utils"
	web "bitbucket.org/envirovisionsolutions/showandtell/web"
)

func Serve() {
	engine := html.NewFileSystem(http.FS(utils.CmdContext.IndexFile), ".html")
	appName := utils.EnvGet("APP_NAME", "Show and Tell")
	isDevelopment := utils.EnvGet("APP_ENV", "production") == "development"
	app := fiber.New(fiber.Config{
		Prefork:           true,
		AppName:           appName,
		Views:             engine,
		EnablePrintRoutes: isDevelopment,
	})

	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC3339,
		Format:     "[${time}] (${ip}) Took: ${latency} ${status} - ${path}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return isDevelopment
		},
	}))

	cacheLengthMinutes, err := strconv.ParseInt(utils.EnvGet("APP_WEB_CACHE_LENGTH_MINUTES", "0"), 10, 32)
	if err != nil {
		cacheLengthMinutes = 0
	}
	if cacheLengthMinutes > 0 {
		log.Printf("INFO: Web route caching active. Timeout (minutes): %d\n", cacheLengthMinutes)
		app.Use(cache.New(cache.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.Query("refresh") == "true"
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
		File:       "web/static/img/favicon.ico",
		FileSystem: http.FS(utils.CmdContext.FavIcon),
	}))

	port := utils.EnvGet("APP_PORT", "8021")
	log.Fatal(app.Listen(":" + port))
}
