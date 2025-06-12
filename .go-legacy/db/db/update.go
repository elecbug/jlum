package db

import (
	"fmt"
)

// UpdateAffiliation updates an affiliation record by a specific field
func UpdateAffiliation(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE affiliation SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdateAuthor updates an author record by a specific field
func UpdateAuthor(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE author SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdateAuthorRole updates an author role record by a specific field
func UpdateAuthorRole(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE author_role SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdateJournalDB updates a journal database record by a specific field
func UpdateJournalDB(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE journal_db SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdateJournal updates a journal record by a specific field
func UpdateJournal(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE journal SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdatePaper updates a paper record by a specific field
func UpdatePaper(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE paper SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdateUser updates a user record by a specific field
func UpdateUser(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE user SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdateRecord updates a record by a specific field
func UpdateRecord(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE record SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdatePaperAuthor updates a paper-author relationship record by a specific field
func UpdatePaperAuthor(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE paper_author SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdateJournalJournalDB updates a journal-database relationship record by a specific field
func UpdateJournalJournalDB(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE journal_journal_db SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdatePaperReference updates a paper reference record by a specific field
func UpdatePaperReference(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE paper_reference SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdatePaperKeyword updates a paper keyword record by a specific field
func UpdatePaperKeyword(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE paper_keyword SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdatePaperTag updates a paper tag record by a specific field
func UpdatePaperTag(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE paper_tag SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}

// UpdatePaperPaperTag updates a paper-tag relationship record by a specific field
func UpdatePaperPaperTag(field string, matchValue interface{}, updateField string, updateValue interface{}) error {
	query := fmt.Sprintf("UPDATE paper_paper_tag SET %s = ? WHERE %s = ?", updateField, field)
	_, err := DB.Exec(query, updateValue, matchValue)
	return err
}
