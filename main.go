package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jgndev/horsefacts-api/api/handler"
	"github.com/jgndev/horsefacts-api/pkg/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	config.PrintConfigStatus()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,OPTIONS",
	}))

	app.Get("/", handler.GetHealthHandler)
	app.Get("/api/facts", handler.GetFactHandler)
	app.Get("/api/breeds", handler.GetBreedHandler)
	app.Get("/api/breeds/:id", handler.GetBreedByIdHandler)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
