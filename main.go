package main

import (
	"fmt"
	"log"

	swagger "github.com/Flussen/swagger-fiber-v3"
	"github.com/NoorUllah43/API-Request-in-Go/db"
	_ "github.com/NoorUllah43/API-Request-in-Go/docs"
	"github.com/NoorUllah43/API-Request-in-Go/routes"

	"github.com/gofiber/fiber/v3/middleware/cors"
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

    app.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"}, 
        AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
    }))

	log.Fatal(app.Listen(":3000"))
}
