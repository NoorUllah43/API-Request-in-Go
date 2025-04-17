package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(useremail string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"useremail": useremail,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenstring, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}
