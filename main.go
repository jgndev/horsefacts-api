package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jgndev/horsefacts-api/api/handler"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	app := fiber.New()

	app.Get("/api/facts", handler.GetFactHandler)
	app.Get("/api/breeds", handler.GetBreedHandler)
	app.Get("/api/breeds/:id", handler.GetBreedByIdHandler)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
