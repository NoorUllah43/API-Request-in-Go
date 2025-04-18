package controllers

import (
	"encoding/json"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/NoorUllah43/API-Request-in-Go/models"
	"github.com/gofiber/fiber/v3"
)

// RegisterUser godoc
// @Summary      Register User
// @Description  Register user using name, email and password, returns JWT token on success.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @scheme		 http
// @Param        credentials  body      models.User  true  "User Name, Email and Password"
// @Success      200
// @Failure      401
// @Router       /auth/registerUser [post]
func RegisterUser(ctx fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return err
	}

	userID, err := db.AddUser(user)
	if err != nil {
		return err
	}
	token, err := middleware.CreateToken(userID)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"token": token,
	}
	return ctx.JSON(response)
}
