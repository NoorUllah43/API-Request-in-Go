package controllers

import (
	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/gofiber/fiber/v3"
)

func Pagenation(ctx fiber.Ctx) error {
	tokenstring := ctx.Get("Authorization")

	if tokenstring == "" {
		return ctx.Status(401).JSON(`missing token`)
	}

	userID, err := middleware.Verifytoken(tokenstring)
	if err != nil {
		return err
	}
	page := ctx.Params("page")

	userData, err := db.GetPagenationData(userID, page)
	if err != nil {
		return err
	}

	return ctx.JSON(userData)
}
