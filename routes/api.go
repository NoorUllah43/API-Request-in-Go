package routes

import (
	"github.com/NoorUllah43/API-Request-in-Go/controllers"
	"github.com/gofiber/fiber/v3"
)

func Routes(app *fiber.App) {

	// api endpoints
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.JSON("welcome to / route")
	})

	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/registerUser", controllers.RegisterUser)
	app.Post("/uploadfile", controllers.Filehandler)

	app.Get("/getUserData", controllers.GetUserData)
	app.Get("/getUserData/:page", controllers.Pagenation)

}
