package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "jlum_user"
	dbName = "jlum_db"
	dbHost = "127.0.0.1"
	dbPort = "3306"
)

var DB *sql.DB

func InitDB(dbPassword string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL with DSN '%s': %v", dsn, err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("DB ping failure: %v", err)
	}
	log.Println("MySQL connnected successfully")
}

// Affiliation represents the affiliation table
type Affiliation struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Author represents the author table
type Author struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	EnglishName   string `json:"english_name"`
	Email         string `json:"email"`
	ORCID         string `json:"orcid"`
	AffiliationID int64  `json:"affiliation_id"`
}

// AuthorRole represents the author_role table
type AuthorRole struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// JournalDB represents the journal_db table
type JournalDB struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Journal represents the journal table
type Journal struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

// Paper represents the paper table
type Paper struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	JournalID int64  `json:"journal_id"`
	Volume    int    `json:"volume"`
	Issue     int    `json:"issue"`
	PageS     int    `json:"page_s"`
	PageE     int    `json:"page_e"`
	Year      int    `json:"year"`
	Month     int    `json:"month"`
	DOI       string `json:"doi"`
	ISBN      string `json:"isbn"`
	ISSN      string `json:"issn"`
	EISSN     string `json:"eissn"`
	Abstract  string `json:"abstract"`
	FileData  []byte `json:"file_data"`
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
}

// PaperAuthor represents the paper_author table
type PaperAuthor struct {
	ID           int64 `json:"id"`
	PaperID      int64 `json:"paper_id"`
	AuthorID     int64 `json:"author_id"`
	AuthorRoleID int64 `json:"author_role_id"`
}

// JournalJournalDB represents the journal_journal_db table
type JournalJournalDB struct {
	ID          int64 `json:"id"`
	JournalID   int64 `json:"journal_id"`
	JournalDBID int64 `json:"journal_db_id"`
}

// PaperReference represents the paper_reference table
type PaperReference struct {
	ID               int64 `json:"id"`
	PaperID          int64 `json:"paper_id"`
	ReferencePaperID int64 `json:"reference_paper_id"`
}

// PaperKeyword represents the paper_keyword table
type PaperKeyword struct {
	ID      int64  `json:"id"`
	PaperID int64  `json:"paper_id"`
	Keyword string `json:"keyword"`
}

// PaperTag represents the paper_tag table
type PaperTag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// PaperPaperTag represents the paper_paper_tag table
type PaperPaperTag struct {
	ID         int64 `json:"id"`
	PaperID    int64 `json:"paper_id"`
	PaperTagID int64 `json:"paper_tag_id"`
}

// User represents the user table
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Record represents the record table
type Record struct {
	ID        int64     `json:"id"`
	PaperID   int64     `json:"paper_id"`
	UserID    int64     `json:"user_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
