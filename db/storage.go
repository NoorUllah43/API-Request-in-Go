package db

import (
	"fmt"

	"github.com/NoorUllah43/API-Request-in-Go/models"
)

func AddUser(userdata models.User) error {
	createUser := fmt.Sprintf(`INSERT INTO
	users (name, email, password)
	VALUES ('%v','%v','%v')`,
		userdata.Name, userdata.Email, userdata.Password,
	)
	DB.QueryRow(createUser)

	return nil
}

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

func InsertResult(data map[string]int) error {
	query := `INSERT INTO 
	result (words, lines, paragraph, punctuation, spaces, vowels, digits, symboles, specialCharacters) 
		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := DB.Exec(query,
		data["words"], data["lines"], data["paragraph"], data["punctuation"],
		data["spaces"], data["vowels"], data["digits"], data["symboles"], data["specialCharacters"],
	)

	if err != nil {
		return err
	}
	return nil
}
