package main

import (
	"fmt"
	"log"

	"github.com/NoorUllah43/API-Request-in-Go/config"
	"github.com/NoorUllah43/API-Request-in-Go/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("loagin env err:",err)
	}

	config.ConnectPostgresqlDB()

	
	routes.Routes(app)
	log.Fatal(app.Listen(":3000"))
}
