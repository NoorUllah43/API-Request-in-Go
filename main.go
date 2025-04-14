package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NoorUllah43/API-Request-in-Go/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()


	err := godotenv.Load()
	
	if err != nil {
		log.Fatal(err)
	}

	routes.Routes(app)

	greet := os.Getenv("GREETING")
	fmt.Println(greet)
	log.Fatal(app.Listen(":3000"))
}
