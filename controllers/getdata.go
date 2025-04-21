package controllers

import (
	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/gofiber/fiber/v3"
)

// GetUserData godoc
// @Summary      Get User Data
// @Description  Retrieves data stored by the user and returns it as JSON.
// @Tags         User
// @Security     ApiKeyAuth
// @Produce      json
// @Success      200 {object} models.ResultData
// @Failure      401 
// @Router       /getUserData [get]
func GetUserData(ctx fiber.Ctx) error {
	tokenstring := ctx.Get("Authorization")

	if tokenstring == "" {
		return ctx.Status(401).JSON(`missing token`)
	}

	userID, err := middleware.Verifytoken(tokenstring)
	if err != nil {
		return err
	}

	userData, err := db.FindUserData(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(userData)
}
