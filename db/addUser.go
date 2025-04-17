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