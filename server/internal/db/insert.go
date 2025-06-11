package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// InsertAffiliation inserts a new affiliation record
func InsertAffiliation(affiliation *Affiliation) (int64, error) {
	res, err := DB.Exec(`INSERT INTO affiliation (name, link) VALUES (?, ?)`,
		affiliation.Name, affiliation.Link)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertAuthor inserts a new author record
func InsertAuthor(author *Author) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO author (name, english_name, email, orcid, affiliation_id)
        VALUES (?, ?, ?, ?, ?)`,
		author.Name, author.EnglishName, author.Email, author.ORCID, author.AffiliationID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertAuthorRole inserts a new author role record
func InsertAuthorRole(role *AuthorRole) (int64, error) {
	res, err := DB.Exec(`INSERT INTO author_role (name) VALUES (?)`, role.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertJournalDB inserts a new journal database record
func InsertJournalDB(journalDB *JournalDB) (int64, error) {
	res, err := DB.Exec(`INSERT INTO journal_db (name) VALUES (?)`, journalDB.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertJournal inserts a new journal record
func InsertJournal(journal *Journal) (int64, error) {
	res, err := DB.Exec(`INSERT INTO journal (title, link) VALUES (?, ?)`,
		journal.Title, journal.Link)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertPaper inserts a new paper record
func InsertPaper(paper *Paper) (int64, error) {
	res, err := DB.Exec(`
        INSERT INTO paper (
            title, journal_id, volume, issue, page_s, page_e, year, month,
            doi, isbn, issn, eissn, abstract, file_data, file_name, file_type
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		paper.Title, paper.JournalID, paper.Volume, paper.Issue,
		paper.PageS, paper.PageE, paper.Year, paper.Month,
		paper.DOI, paper.ISBN, paper.ISSN, paper.EISSN,
		paper.Abstract, paper.FileData, paper.FileName, paper.FileType)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertPaperAuthor inserts a new paper-author relationship record
func InsertPaperAuthor(paperAuthor *PaperAuthor) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_author (paper_id, author_id, author_role_id) VALUES (?, ?, ?)`,
		paperAuthor.PaperID, paperAuthor.AuthorID, paperAuthor.AuthorRoleID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertJournalJournalDB inserts a new journal-database relationship record
func InsertJournalJournalDB(link *JournalJournalDB) (int64, error) {
	res, err := DB.Exec(`INSERT INTO journal_journal_db (journal_id, journal_db_id) VALUES (?, ?)`,
		link.JournalID, link.JournalDBID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertPaperReference inserts a new paper reference record
func InsertPaperReference(reference *PaperReference) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_reference (paper_id, reference_paper_id) VALUES (?, ?)`,
		reference.PaperID, reference.ReferencePaperID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertPaperKeyword inserts a new paper keyword record
func InsertPaperKeyword(keyword *PaperKeyword) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_keyword (paper_id, keyword) VALUES (?, ?)`,
		keyword.PaperID, keyword.Keyword)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertPaperTag inserts a new paper tag record
func InsertPaperTag(tag *PaperTag) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_tag (name) VALUES (?)`, tag.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertPaperPaperTag inserts a new paper-tag relationship record
func InsertPaperPaperTag(link *PaperPaperTag) (int64, error) {
	res, err := DB.Exec(`INSERT INTO paper_paper_tag (paper_id, paper_tag_id) VALUES (?, ?)`,
		link.PaperID, link.PaperTagID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertUser inserts a new user record
func InsertUser(user *User) (int64, error) {
	res, err := DB.Exec(`INSERT INTO user (username, password, email) VALUES (?, ?, ?)`,
		user.Username, user.Password, user.Email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// InsertRecord inserts a new record
func InsertRecord(record *Record) (int64, error) {
	now := time.Now()
	res, err := DB.Exec(`
        INSERT INTO record (paper_id, user_id, note, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)`,
		record.PaperID, record.UserID, record.Note, now, now)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
