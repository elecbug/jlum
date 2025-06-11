package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// affiliation
func InsertAffiliation(name, link string) (int64, error) {
	res, err := DB.Exec(`INSERT INTO affiliation (name, link) VALUES (?, ?)`, name, link)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// author
func InsertAuthor(name, englishName, email, orcid string, affiliationID int64) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO author (name, english_name, email, orcid, affiliation_id)
        VALUES (?, ?, ?, ?, ?)`, name, englishName, email, orcid, affiliationID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// author_role
func InsertAuthorRole(name string) (int64, error) {
	res, err := DB.Exec(`INSERT INTO author_role (name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// journal_db
func InsertJournalDB(name string) (int64, error) {
	res, err := DB.Exec(`INSERT INTO journal_db (name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// journal
func InsertJournal(title, link string) (int64, error) {
	res, err := DB.Exec(`INSERT INTO journal (title, link) VALUES (?, ?)`, title, link)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// paper
func InsertPaper(title string, journalID, volume, issue, pageS, pageE, year, month int,
	doi, isbn, issn, eissn, abstract, fileName, fileType string, fileData []byte) (int64, error) {

	res, err := DB.Exec(`
        INSERT INTO paper (
            title, journal_id, volume, issue, page_s, page_e, year, month,
            doi, isbn, issn, eissn, abstract, file_data, file_name, file_type
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		title, journalID, volume, issue, pageS, pageE, year, month,
		doi, isbn, issn, eissn, abstract, fileData, fileName, fileType)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// paper_author
func InsertPaperAuthor(paperID, authorID, roleID int64) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_author (paper_id, author_id, author_role_id) VALUES (?, ?, ?)`,
		paperID, authorID, roleID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// journal_journal_db
func InsertJournalDBLink(journalID, dbID int64) (int64, error) {
	res, err := DB.Exec(`INSERT INTO journal_journal_db (journal_id, journal_db_id) VALUES (?, ?)`, journalID, dbID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// paper_reference
func InsertPaperReference(paperID, refID int64) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_reference (paper_id, reference_paper_id) VALUES (?, ?)`, paperID, refID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// paper_keyword
func InsertPaperKeyword(paperID int64, keyword string) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_keyword (paper_id, keyword) VALUES (?, ?)`, paperID, keyword)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// paper_tag
func InsertPaperTag(name string) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_tag (name) VALUES (?)`, name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// paper_paper_tag
func LinkPaperTag(paperID, tagID int64) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_paper_tag (paper_id, paper_tag_id) VALUES (?, ?)`, paperID, tagID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// user
func InsertUser(username, password, email string) (int64, error) {
	res, err := DB.Exec(`INSERT INTO user (username, password, email) VALUES (?, ?, ?)`, username, password, email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// record
func InsertRecord(paperID, userID int64, note string) (int64, error) {
	now := time.Now()
	res, err := DB.Exec(`
        INSERT INTO record (paper_id, user_id, note, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)`, paperID, userID, note, now, now)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
