package controllers

import (
	"fmt"
	"os"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/gofiber/fiber/v3"
)

// Filehandler godoc
// @Summary      Upload a file
// @Description  Upload a file to analyze, return json
// @Tags         User
// @Accept       multipart/form-data
// @Produce      json
// @host 		 localhost:3000
// @BasePath 	 /
// @scheme		 http
// @Success      200
// @Failure      401
// @Router       /uploadfile [post]
func Filehandler(ctx fiber.Ctx) error {

	tokenstring := ctx.Get("Authorization")

	if tokenstring == "" {
		return ctx.Status(401).JSON(`statuscode : 401, missing token`)
	}

	userID, err := middleware.Verifytoken(tokenstring)
	if err != nil {
		return err
	}
// 
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	fileLocation := fmt.Sprintf("./uploads/%v", file.Filename)

	err = ctx.SaveFile(file, fileLocation)
	if err != nil {
		return err
	}

	readfile, err := os.ReadFile("./uploads/file.txt")
	if err != nil {
		panic(err)
	}

	fileContent := string(readfile)

	data := analyze(fileContent)

	err = db.InsertResult(data, userID)
	if err != nil {
		return err
	}

	return ctx.JSON(data)

}
