package db

import (
	"fmt"
	"strconv"

	"github.com/NoorUllah43/API-Request-in-Go/models"
)

func GetPagenationData(userID int, limitdata string) ([]models.ResultData, error) {

	pagenum, err := strconv.Atoi(limitdata)
	if err != nil {
		return []models.ResultData{}, err
	}

	quray := fmt.Sprintf(`SELECT 
	words, 
	vowels, 
	spaces,
	punctuation,
	paragraph,
	specialcharacters,
	symboles,
	digits,
	lines FROM result WHERE id=$1 LIMIT %v OFFSET 0`, pagenum)

	var userData []models.ResultData

	rows, err := DB.Query(quray, userID)
	if err != nil {
		return userData, err
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
