package routes

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
)

func Routes(app *fiber.App) {

	app.Post("/uploadfile", filehandler)

}




func filehandler(ctx fiber.Ctx) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	fileLocation := fmt.Sprintf("./uploads/%v", file.Filename)

	if err := ctx.SaveFile(file, fileLocation); err != nil {
		return err
	}

	
	readfile, ferr := os.ReadFile("./uploads/file.txt")
	if ferr != nil {
		panic(ferr)
	}
	
	fileContent := string(readfile)






	return ctx.JSON(fileContent)
}