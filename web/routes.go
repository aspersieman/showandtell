package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	auth "bitbucket.org/envirovisionsolutions/showandtell/auth"
)

func SetupRoutes(app *fiber.App, keycloak *auth.Keycloak) {
	controller := NewController(keycloak)

	app.Get("/", controller.Home)

	// Web UI routes
	// TODO: Find cleaner way to handle client routes
	app.Get("/schedule/*", controller.Home)
	app.Get("/schedules/*", controller.Home)
	app.Get("/login/*", controller.Home)

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/auth/login", controller.ApiAuthLogin)
	v1.Post("/auth/logout", controller.ApiAuthLogout)
	v1.Post("/auth/token", controller.ApiAuthToken)
	v1.Get("/schedules/:id<int>?", controller.ApiGetSchedules)
	v1.Get("/speakers/:id<int>?", controller.ApiGetSpeakers)
	v1.Get("/version", controller.ApiGetVersion)

	keycloakMiddleware := NewMiddleware(keycloak)
	v1Auth := api.Group("/v1", keycloakMiddleware.VerifyToken)
	v1Auth.Get("/users/:id<int>?", controller.ApiGetUsers)
	v1Auth.Post("/schedules", controller.ApiPostSchedules)
	v1Auth.Put("/schedules/:id<int>", controller.ApiPutSchedules)
	v1Auth.Delete("/schedules/:id<int>", controller.ApiDeleteSchedules)
	v1Auth.Post("/speakers", controller.ApiPostSpeakers)
	v1Auth.Put("/speakers/:id<int>", controller.ApiPutSpeakers)
	v1Auth.Delete("/speakers/:id<int>", controller.ApiDeleteSpeakers)
}
