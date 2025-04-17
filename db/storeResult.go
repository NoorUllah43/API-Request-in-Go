package db


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