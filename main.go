package main

import (
	"log"
	"github.com/NoorUllah43/API-Request-in-Go/routes"
	"github.com/gofiber/fiber/v3"
	
)

func main() {
	app := fiber.New()

	
	routes.Routes(app)
	log.Fatal(app.Listen(":3000"))
}
