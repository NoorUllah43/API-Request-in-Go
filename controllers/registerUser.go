package controllers

import (
	"encoding/json"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/NoorUllah43/API-Request-in-Go/models"
	"github.com/gofiber/fiber/v3"
)

// @Summary		Register user
// @Description	Using this route user will register in the database
// @Produce		json
// @Success		200
// @Router		/auth/registerUser [post]
func RegisterUser(ctx fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return err
	}

	err = db.AddUser(user)
	if err != nil {
		return err
	}
	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		return err
	}
	return ctx.JSON(token)
}
