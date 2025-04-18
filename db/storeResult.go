package db


func InsertResult(data map[string]int, userID int) error {
	query := `INSERT INTO 
	result (id,words, vowels, spaces, punctuation, paragraph, specialcharacters, symboles, digits, lines) 
		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := DB.Exec(query,
		userID,data["words"], data["vowels"], data["spaces"], data["punctuation"], data["paragraph"], data["specialCharacters"], data["symboles"], data["digits"], data["lines"],
	)

	if err != nil {
		return err
	}
	return nil
}