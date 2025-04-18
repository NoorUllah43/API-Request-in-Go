package db

import "fmt"

func FindUserData(userID int) error {

	quray := `SELECT * FROM result WHERE id=$1`

	rows, err := DB.Query(quray, userID)

	if err != nil {
		return err
	}

	fmt.Println(rows)

	return nil
}