package controllers

import (
	"encoding/json"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/NoorUllah43/API-Request-in-Go/models"
	"github.com/gofiber/fiber/v3"
)




func Login(ctx fiber.Ctx) error {
	var credentials models.UserCredentials
	var email string
	var password string

	err := json.Unmarshal(ctx.Body(), &credentials)
	if err != nil {
		return err
	}


	email, password, err = db.FindUser(credentials)
	if err != nil {
		
		return ctx.JSON(`Unautherize please provide correct email and password`)
	}
	if email != credentials.Email && password != credentials.Password {
		return ctx.JSON(`unautherize please provide correct email and password`)
	}

	token, err := middleware.CreateToken(credentials.Email)
	if err != nil {
		return err
	}
	
	return ctx.JSON(token)

}