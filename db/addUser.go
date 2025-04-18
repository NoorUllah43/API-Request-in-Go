package db

import (
	"github.com/NoorUllah43/API-Request-in-Go/models"
)

func AddUser(userdata models.User) (int, error) {

	query :=
		`INSERT INTO
	  users (name, email, password)
	VALUES
	  ($1, $2, $3)
	RETURNING id`

	stmt, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var userID int

	err = stmt.QueryRow(userdata.Name, userdata.Email, userdata.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
