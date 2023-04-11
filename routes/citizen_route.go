package routes

import (
	"github.com/TheBigSteph/fullstack-test/controllers"
	"github.com/gofiber/fiber/v2"
)


func CitizenRoute(app *fiber.App) {
	app.Post("/citizens", controllers.CreateCitizen)
}
