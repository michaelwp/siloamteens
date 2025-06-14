package router

import "github.com/gofiber/fiber/v3"

func Router(app *fiber.App) {
	// @Summary Healthcheck
	// @Description to check service health condition, if healthy will return "Hello, World 👋!"
	// @Produce  text
	// @Success 200 {object} User
	// @Router /healthcheck [get]
	app.Get("/healthcheck", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	api := app.Group("/api")
	routerV1(api)
}

func routerV1(api fiber.Router) {
	v1 := api.Group("/v1")
	v1.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, V1!")
	})
}
