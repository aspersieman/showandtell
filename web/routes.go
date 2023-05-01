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
	v1.Post("/auth/logout", controller.ApiAuthLogout)
	v1.Get("/schedules/:id<int>?", controller.ApiGetSchedules)
	v1.Get("/speakers/:id<int>?", controller.ApiGetSpeakers)
	v1.Get("/version", controller.ApiGetVersion)

	// TODO: The following routes should be protected by auth
	v1.Post("/schedules", controller.ApiPostSchedules)
	v1.Put("/schedules/:id<int>", controller.ApiPutSchedules)
	v1.Delete("/schedules/:id<int>", controller.ApiDeleteSchedules)
	v1.Post("/speakers", controller.ApiPostSpeakers)
	v1.Put("/speakers/:id<int>", controller.ApiPutSpeakers)
	v1.Delete("/speakers/:id<int>", controller.ApiDeleteSpeakers)

	keycloakMiddleware := NewMiddleware(keycloak)
	v1Auth := api.Group("/v1", keycloakMiddleware.VerifyToken)
	v1Auth.Get("/users/:id<int>?", controller.ApiGetUsers)
}
