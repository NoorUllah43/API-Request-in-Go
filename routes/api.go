package routes

import (
	"github.com/NoorUllah43/API-Request-in-Go/controllers"
	"github.com/gofiber/fiber/v3"
)

// api endpoint
func Routes(app *fiber.App) {

	app.Post("/uploadfile", controllers.Filehandler)
	app.Post("auth/login", controllers.Login)

}












