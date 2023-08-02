package main

import (
	"fmt"
	"log"

	"github.com/anuj0809/Backend_AS/controllers"
	"github.com/anuj0809/Backend_AS/database"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello Wrold")
}

func main() {
	database.ConnectToDB()
	// create a new fiber instance
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		// Logging the request information
		fmt.Printf("Request: %s %s \n", c.Method(), c.Path())

		// Proceed with the next middleware/handler
		return c.Next()
	})

	app.Get("/api", welcome)
	app.Post("players", controllers.CratePlayer)
	app.Get("/players", controllers.GetAllPlayers)
	app.Get("/players/random", controllers.GetRandomPlayer)
	app.Get("/players/rank/:val", controllers.GetPlayerByRank)
	app.Put("/players/:id", controllers.UpdatePlayer)
	app.Delete("/players/:id", controllers.DeletePlayer)
	log.Fatal(app.Listen(":3000"))
}
