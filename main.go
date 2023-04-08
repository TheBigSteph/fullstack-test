package main

import (
	"github.com/TheBigSteph/fullstack-test/configs"
	"github.com/TheBigSteph/fullstack-test/routes"
	"github.com/gofiber/fiber/v2"
)

func main () {
	app := fiber.New()

	configs.ConnectDB()

	routes.AdministratorRoute(app)
	routes.CitizenRoute(app)
	routes.RequestRoute(app)

	app.Listen(":8080")
}
