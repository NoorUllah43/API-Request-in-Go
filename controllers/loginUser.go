package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/NoorUllah43/API-Request-in-Go/models"
	"github.com/gofiber/fiber/v3"
)

// Login godoc
// @Summary      Login User
// @Description  Logs in a user using email and password, returns JWT token on success.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @host 		 localhost:3000
// @BasePath 	 /
// @scheme		 http
// @Param        credentials  body      models.UserCredentials  true  "User Email and Password"
// @Success      200
// @Failure      401
// @Router       /auth/login [post]
func Login(ctx fiber.Ctx) error {
	var credentials models.UserCredentials
	err := json.Unmarshal(ctx.Body(), &credentials)
	if err != nil {
		return err
	}

	email, password, userID, err := db.FindUser(credentials)
	if err != nil {
		return ctx.JSON(`unautherize please provide correct email and password`)
	}
	
	if email != credentials.Email && password != credentials.Password {
		return ctx.JSON(`unautherize please provide correct email and password`)
	}
	fmt.Println(userID)

	token, err := middleware.CreateToken(userID)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"token": token,
	}

	return ctx.JSON(response)

}
