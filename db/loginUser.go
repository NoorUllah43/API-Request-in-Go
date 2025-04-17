package db

import "github.com/NoorUllah43/API-Request-in-Go/models"

func FindUser(userCredentials models.UserCredentials) (string, string, error) {
	var email string
	var password string

	user := `SELECT email, password FROM users WHERE email=$1`

	err := DB.QueryRow(user, userCredentials.Email).Scan(&email, &password)
	if err != nil {
		return "", "", err
	}

	return email, password, nil
}