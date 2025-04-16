package controllers

import (
	"encoding/json"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/NoorUllah43/API-Request-in-Go/models"
	"github.com/gofiber/fiber/v3"
)

func RegisterUser(ctx fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return err
	}
	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		return err
	}

	db.AddUser(user)

	return ctx.JSON(token)
}