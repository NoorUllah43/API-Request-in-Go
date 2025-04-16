package db

import (
	"fmt"

	"github.com/NoorUllah43/API-Request-in-Go/config"
	"github.com/NoorUllah43/API-Request-in-Go/models"
)

func InsertUser(userdata models.User) error {
	createUser := fmt.Sprintf(`INSERT INTO
	users (name, email, password)
	VALUES ('%v','%v','%v')`,
		userdata.Name, userdata.Email, userdata.Password,
	)
	config.DB.QueryRow(createUser)
	return nil
}

func FindUser(userCredentials models.UserCredentials) error {

	

	return nil
}
