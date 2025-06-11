package db

import (
	"fmt"
)

// GetAffiliation retrieves an affiliation record by a specific field
func GetAffiliation(field string, value interface{}) (*Affiliation, error) {
	query := fmt.Sprintf("SELECT id, name, link FROM affiliation WHERE %s = ?", field)
	row := DB.QueryRow(query, value)
	affiliation := &Affiliation{}
	if err := row.Scan(&affiliation.ID, &affiliation.Name, &affiliation.Link); err != nil {
		return nil, err
	}
	return affiliation, nil
}

// GetAuthor retrieves an author record by a specific field
func GetAuthor(field string, value interface{}) (*Author, error) {
	query := fmt.Sprintf("SELECT id, name, english_name, email, orcid, affiliation_id FROM author WHERE %s = ?", field)
	row := DB.QueryRow(query, value)
	author := &Author{}
	if err := row.Scan(&author.ID, &author.Name, &author.EnglishName, &author.Email, &author.ORCID, &author.AffiliationID); err != nil {
		return nil, err
	}
	return author, nil
}

// GetAuthorRole retrieves an author role record by a specific field
func GetAuthorRole(field string, value interface{}) (*AuthorRole, error) {
	query := fmt.Sprintf("SELECT id, name FROM author_role WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	authorRole := &AuthorRole{}
	if err := row.Scan(&authorRole.ID, &authorRole.Name); err != nil {
		return nil, err
	}
	return authorRole, nil
}

// GetJournalDB retrieves a journal database record by a specific field
func GetJournalDB(field string, value interface{}) (*JournalDB, error) {
	query := fmt.Sprintf("SELECT id, name FROM journal_db WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	journalDB := &JournalDB{}
	if err := row.Scan(&journalDB.ID, &journalDB.Name); err != nil {
		return nil, err
	}
	return journalDB, nil
}

// GetJournal retrieves a journal record by a specific field
func GetJournal(field string, value interface{}) (*Journal, error) {
	query := fmt.Sprintf("SELECT id, title, link FROM journal WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	journal := &Journal{}
	if err := row.Scan(&journal.ID, &journal.Title, &journal.Link); err != nil {
		return nil, err
	}
	return journal, nil
}

// GetPaper retrieves a paper record by a specific field
func GetPaper(field string, value interface{}) (*Paper, error) {
	query := fmt.Sprintf(`SELECT id, title, journal_id, volume, issue, page_s, page_e, year, month, doi, isbn, issn, eissn, abstract, file_data, file_name, file_type FROM paper WHERE %s = ?`, field)
	row := DB.QueryRow(query, value)

	paper := &Paper{}
	if err := row.Scan(&paper.ID, &paper.Title, &paper.JournalID, &paper.Volume, &paper.Issue, &paper.PageS, &paper.PageE, &paper.Year, &paper.Month, &paper.DOI, &paper.ISBN, &paper.ISSN, &paper.EISSN, &paper.Abstract, &paper.FileData, &paper.FileName, &paper.FileType); err != nil {
		return nil, err
	}
	return paper, nil
}

// GetUser retrieves a user record by a specific field
func GetUser(field string, value interface{}) (*User, error) {
	query := fmt.Sprintf("SELECT id, username, password, email FROM user WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	user := &User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
		return nil, err
	}
	return user, nil
}

// GetRecord retrieves a record by a specific field
func GetRecord(field string, value interface{}) (*Record, error) {
	query := fmt.Sprintf("SELECT id, paper_id, user_id, note, created_at, updated_at FROM record WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	record := &Record{}
	if err := row.Scan(&record.ID, &record.PaperID, &record.UserID, &record.Note, &record.CreatedAt, &record.UpdatedAt); err != nil {
		return nil, err
	}
	return record, nil
}

// GetPaperAuthor retrieves a paper-author relationship record by a specific field
func GetPaperAuthor(field string, value interface{}) (*PaperAuthor, error) {
	query := fmt.Sprintf("SELECT id, paper_id, author_id, author_role_id FROM paper_author WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	paperAuthor := &PaperAuthor{}
	if err := row.Scan(&paperAuthor.ID, &paperAuthor.PaperID, &paperAuthor.AuthorID, &paperAuthor.AuthorRoleID); err != nil {
		return nil, err
	}
	return paperAuthor, nil
}

// GetJournalJournalDB retrieves a journal-database relationship record by a specific field
func GetJournalJournalDB(field string, value interface{}) (*JournalJournalDB, error) {
	query := fmt.Sprintf("SELECT id, journal_id, journal_db_id FROM journal_journal_db WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	journalJournalDB := &JournalJournalDB{}
	if err := row.Scan(&journalJournalDB.ID, &journalJournalDB.JournalID, &journalJournalDB.JournalDBID); err != nil {
		return nil, err
	}
	return journalJournalDB, nil
}

// GetPaperReference retrieves a paper reference record by a specific field
func GetPaperReference(field string, value interface{}) (*PaperReference, error) {
	query := fmt.Sprintf("SELECT id, paper_id, reference_paper_id FROM paper_reference WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	paperReference := &PaperReference{}
	if err := row.Scan(&paperReference.ID, &paperReference.PaperID, &paperReference.ReferencePaperID); err != nil {
		return nil, err
	}
	return paperReference, nil
}

// GetPaperKeyword retrieves a paper keyword record by a specific field
func GetPaperKeyword(field string, value interface{}) (*PaperKeyword, error) {
	query := fmt.Sprintf("SELECT id, paper_id, keyword FROM paper_keyword WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	paperKeyword := &PaperKeyword{}
	if err := row.Scan(&paperKeyword.ID, &paperKeyword.PaperID, &paperKeyword.Keyword); err != nil {
		return nil, err
	}
	return paperKeyword, nil
}

// GetPaperTag retrieves a paper tag record by a specific field
func GetPaperTag(field string, value interface{}) (*PaperTag, error) {
	query := fmt.Sprintf("SELECT id, name FROM paper_tag WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	paperTag := &PaperTag{}
	if err := row.Scan(&paperTag.ID, &paperTag.Name); err != nil {
		return nil, err
	}
	return paperTag, nil
}

// GetPaperPaperTag retrieves a paper-tag relationship record by a specific field
func GetPaperPaperTag(field string, value interface{}) (*PaperPaperTag, error) {
	query := fmt.Sprintf("SELECT id, paper_id, paper_tag_id FROM paper_paper_tag WHERE %s = ?", field)
	row := DB.QueryRow(query, value)

	paperPaperTag := &PaperPaperTag{}
	if err := row.Scan(&paperPaperTag.ID, &paperPaperTag.PaperID, &paperPaperTag.PaperTagID); err != nil {
		return nil, err
	}
	return paperPaperTag, nil
}
