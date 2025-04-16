package controllers

import (
	"fmt"
	"os"

	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/gofiber/fiber/v3"
)

func Filehandler(ctx fiber.Ctx) error {

	tokenstring := ctx.Get("Authorization")

	if tokenstring == "" {
		return ctx.JSON(`statuscode : 401, missing token`)
	}

	err := middleware.Verifytoken(tokenstring)
	if err != nil {
		return err
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	fileLocation := fmt.Sprintf("./uploads/%v", file.Filename)

	if err := ctx.SaveFile(file, fileLocation); err != nil {
		return err
	}

	readfile, err := os.ReadFile("./uploads/file.txt")
	if err != nil {
		panic(err)
	}

	fileContent := string(readfile)

	data := analyze(fileContent)

	return ctx.JSON(data)

}