package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtsecretkey = []byte("thisisjwtkey")

func CreateToken(useremail string) (string, error) {

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

func Verifytoken(tokenstring string) error {
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
