package controllers

import (
	"fmt"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/gofiber/fiber/v3"
)

func GetUserData(ctx fiber.Ctx) error {
	tokenstring := ctx.Get("Authorization")

	if tokenstring == "" {
		return ctx.JSON(`statuscode : 401, missing token`)
	}

	userID, err := middleware.Verifytoken(tokenstring)
	if err != nil {
		return err
	}

	fmt.Println(userID)
	db.FindUserData(userID)

	return nil
}
