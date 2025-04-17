package main

import (
	"fmt"
	"log"

	swagger "github.com/Flussen/swagger-fiber-v3"
	"github.com/NoorUllah43/API-Request-in-Go/db"
	_ "github.com/NoorUllah43/API-Request-in-Go/docs"
	"github.com/NoorUllah43/API-Request-in-Go/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

// @title			Golang Task
// @version		    3.0
// @host			localhost:3000
// @BasePath		/
func main() {

	app := fiber.New()
	app.Get("/swagger/*", swagger.New())

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("loagin env err:", err)
	}

	db.ConnectPostgresqlDB()

	routes.Routes(app)
	log.Fatal(app.Listen(":3000"))
}
