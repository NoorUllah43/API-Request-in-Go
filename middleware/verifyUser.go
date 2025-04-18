package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Verifytoken(tokenstring string) (int, error) {

	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	
	floatID, ok := claims["userID"].(float64)
	if !ok {
		return 0, fmt.Errorf("userID doesn't exist")
	}

	userID := int(floatID)

	return userID, nil
}
