package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NoorUllah43/API-Request-in-Go/db"
	"github.com/NoorUllah43/API-Request-in-Go/middleware"
	"github.com/NoorUllah43/API-Request-in-Go/models"
	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
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

func RegisterUser(ctx fiber.Ctx) error {
	var user models.User

	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return err
	}
	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		return err
	}

	db.InsertUser(user)

	return ctx.JSON(token)
}

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

func analyze(str string) map[string]int {

	chunkeData := make(map[string]int)

	for i := 0; i < len(str); i++ {
		chunkeData["character"]++
		if str[i] == (' ') {
			chunkeData["spaces"]++
		}
		if str[i] == ('.') {
			chunkeData["lines"]++
		}
		if str[i] == ('\n') {
			chunkeData["paragraph"]++
		}
		if str[i] == (' ') || str[i] == ('\n') {
			chunkeData["words"]++
		}
		if str[i] == (';') || str[i] == (':') || str[i] == ('\'') || str[i] == ('"') || str[i] == (',') || str[i] == ('.') || str[i] == ('?') {
			chunkeData["punctuation"]++
		}
		if str[i] == ('[') || str[i] == (']') || str[i] == ('{') || str[i] == ('}') || str[i] == ('(') || str[i] == (')') || str[i] == ('|') || str[i] == ('\\') || str[i] == ('/') || str[i] == ('+') || str[i] == ('=') || str[i] == ('<') || str[i] == ('>') {
			chunkeData["symboles"]++
		}
		if str[i] == ('!') || str[i] == ('@') || str[i] == ('#') || str[i] == ('$') || str[i] == ('%') || str[i] == ('^') || str[i] == ('&') || str[i] == ('*') || str[i] == ('~') {
			chunkeData["specialCharacters"]++
		}
		if str[i] == ('0') || str[i] == ('1') || str[i] == ('2') || str[i] == ('3') || str[i] == ('4') || str[i] == ('5') || str[i] == ('6') || str[i] == ('7') || str[i] == ('8') || str[i] == ('9') {
			chunkeData["digits"]++
		}
		if str[i] == ('a') || str[i] == ('e') || str[i] == ('i') || str[i] == ('o') || str[i] == ('u') {
			chunkeData["vowels"]++
		}
	}

	return chunkeData

}
