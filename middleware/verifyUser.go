package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Verifytoken(tokenstring string) error {

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["useremail"])

	return nil
}
