package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"github.com/NoorUllah43/API-Request-in-Go/models"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Filehandler(ctx fiber.Ctx) error {

	tokenstring := ctx.Get("Authorization")

	if tokenstring == "" {
		return ctx.JSON(`statuscode : 401, missing token`)
	}
	

	err := verifytoken(tokenstring)
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



func Login(ctx fiber.Ctx) error {
	var user models.Person

	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return err
	}
	token, err := createToken(user.Email)
	if err != nil {
		return err
	}

	return ctx.JSON(token) 
}





var jwtsecretkey = []byte("thisisjwtkey")

func createToken(useremail string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"useremail": useremail,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenstring, err := token.SignedString(jwtsecretkey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}

func verifytoken(tokenstring string) error {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return jwtsecretkey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
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
