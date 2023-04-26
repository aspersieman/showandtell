package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	auth "bitbucket.org/envirovisionsolutions/showandtell/auth"
)

func SetupRoutes(app *fiber.App, keycloak *auth.Keycloak) {
	controller := NewController(keycloak)
	app.Get("/", controller.Home)
	app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/auth/login", controller.ApiAuthLogin)
	v1.Get("/schedules/:id<int>?", controller.ApiGetSchedules)

	keycloakMiddleware := NewMiddleware(keycloak)
	v1Auth := api.Group("/v1", keycloakMiddleware.VerifyToken)
	v1Auth.Get("/users/:id<int>?", controller.ApiGetUsers)
	v1Auth.Get("/version", controller.ApiGetVersion)
}
