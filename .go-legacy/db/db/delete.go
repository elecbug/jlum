package db

import (
	"fmt"
)

// DeleteAffiliation deletes an affiliation record by a specific field
func DeleteAffiliation(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM affiliation WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeleteAuthor deletes an author record by a specific field
func DeleteAuthor(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM author WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeleteAuthorRole deletes an author role record by a specific field
func DeleteAuthorRole(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM author_role WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeleteJournalDB deletes a journal database record by a specific field
func DeleteJournalDB(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM journal_db WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeleteJournal deletes a journal record by a specific field
func DeleteJournal(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM journal WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeletePaper deletes a paper record by a specific field
func DeletePaper(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM paper WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeleteUser deletes a user record by a specific field
func DeleteUser(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM user WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeleteRecord deletes a record by a specific field
func DeleteRecord(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM record WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeletePaperAuthor deletes a paper-author relationship record by a specific field
func DeletePaperAuthor(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM paper_author WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeleteJournalJournalDB deletes a journal-database relationship record by a specific field
func DeleteJournalJournalDB(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM journal_journal_db WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeletePaperReference deletes a paper reference record by a specific field
func DeletePaperReference(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM paper_reference WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeletePaperKeyword deletes a paper keyword record by a specific field
func DeletePaperKeyword(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM paper_keyword WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeletePaperTag deletes a paper tag record by a specific field
func DeletePaperTag(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM paper_tag WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}

// DeletePaperPaperTag deletes a paper-tag relationship record by a specific field
func DeletePaperPaperTag(field string, value interface{}) error {
	query := fmt.Sprintf("DELETE FROM paper_paper_tag WHERE %s = ?", field)
	_, err := DB.Exec(query, value)
	return err
}
