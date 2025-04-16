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

	err := DB.QueryRow(user, userCredentials.Email ).Scan(&email, &password)
	if err != nil {
		return "", "", err
	}

	return email, password, nil
}
