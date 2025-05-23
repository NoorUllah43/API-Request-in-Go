package db

import (

	"github.com/NoorUllah43/API-Request-in-Go/models"
)

func FindUserData(userID int) ([]models.ResultData, error) {
	
	

	quray := `SELECT 
	words, 
	vowels, 
	spaces,
	punctuation,
	paragraph,
	specialcharacters,
	symboles,
	digits,
	lines FROM result WHERE id=$1`

	
	var userData []models.ResultData

	rows, err := DB.Query(quray, userID)
	if err != nil {
		return userData , err
	}

	defer rows.Close()
	for rows.Next() {
		var result models.ResultData
		err := rows.Scan(&result.Words, &result.Vowels, &result.Spaces, &result.Punctuation, &result.Paragraph, &result.Specialcharacters, &result.Symboles, &result.Digits, &result.Lines)
		if err != nil {
			return userData, err
		}
		userData = append(userData, result)
		
	}

	

	return userData, nil
}
