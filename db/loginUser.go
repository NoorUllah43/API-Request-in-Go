package db

import "github.com/NoorUllah43/API-Request-in-Go/models"

func FindUser(userCredentials models.UserCredentials) (string, string, int, error) {
	var email string
	var password string
	var id int

	user := `SELECT email, password, id FROM users WHERE email=$1`

	err := DB.QueryRow(user, userCredentials.Email).Scan(&email, &password, &id)
	if err != nil {
		return "", "", 0, err
	}

	return email, password, id, nil
}