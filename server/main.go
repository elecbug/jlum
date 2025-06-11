package main

import (
	"flag"
	"jlum-server/internal/db"
)

func main() {
	dbPassword := flag.String("db-password", "", "MySQL password for jlum_user")
	flag.Parse()

	db.InitDB(*dbPassword)
	defer db.DB.Close()

	// 사용 예시
	affiliationID, _ := db.InsertAffiliation("KNU", "https://knu.ac.kr")
	authorRoleID, _ := db.InsertAuthorRole("제1저자")
	journalID, _ := db.InsertJournal("Journal of AI", "https://ai-journal.org")
	journalDbID, _ := db.InsertJournalDB("Scopus")
	paperTagID, _ := db.InsertPaperTag("AI")

	authorID, _ := db.InsertAuthor("홍길동", "Gil-Dong Hong", "gil@example.com", "0000-0000-0000-0000", affiliationID)
	paperID, _ := db.InsertPaper("논문 제목", int(journalID), 1, 1, 100, 120, 2025, 6,
		"10.doi/abc", "isbn", "issn", "eissn", "초록입니다", "test name", "", ([]byte)("temp"))

	db.InsertPaperAuthor(paperID, authorID, authorRoleID)
	db.InsertJournalDBLink(journalID, journalDbID)
	db.InsertPaperKeyword(paperID, "딥러닝")
	db.LinkPaperTag(paperID, paperTagID)

	userID, _ := db.InsertUser("admin", "hashed_password", "admin@example.com")
	db.InsertRecord(paperID, userID, "중요 논문입니다.")
}
