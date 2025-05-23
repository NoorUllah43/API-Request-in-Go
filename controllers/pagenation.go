package controllers

import (
	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/gofiber/fiber/v3"
)


// Pagenation godoc
// @Summary      Get User Data (Paginated)
// @Description  Retrieves user data by page number. Returns paginated JSON data.
// @Tags         User
// @Security     ApiKeyAuth
// @Param        page path int true "Page number"
// @Produce      json
// @Success      200 {array}     models.ResultData
// @Failure      401 {string}    "Unauthorized"
// @Router       /getUserData/{page} [get]
func Pagenation(ctx fiber.Ctx) error {
	tokenstring := ctx.Get("Authorization")

	if tokenstring == "" {
		return ctx.Status(401).JSON(`Unauthorized`)
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
