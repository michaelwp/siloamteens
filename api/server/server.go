package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/michaelwp/siloamteens/api/router"
)

func Server(port string) error {
	app := fiber.New()
	router.Router(app)
	return app.Listen(port)
}
