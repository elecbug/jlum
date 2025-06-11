package paper

import "time"

type Paper struct {
	ID           int      `json:"id"`
	AuthorIDs    []int    `json:"authors"`
	Title        string   `json:"title"`
	JournalID    int      `json:"journal"`
	Volume       int      `json:"volume"`
	Issue        int      `json:"issue"`
	PageS        int      `json:"page_s"`
	PageE        int      `json:"page_e"`
	Year         int      `json:"year"`
	Month        int      `json:"month"`
	DOI          string   `json:"doi"`
	ISBN         string   `json:"isbn"`
	ISSN         string   `json:"issn"`
	EISSN        string   `json:"eissn"`
	Abstract     string   `json:"abstract"`
	FileData     []byte   `json:"file_data"`
	FileName     string   `json:"file_name"`
	FileType     string   `json:"file_type"`
	Keywords     []string `json:"keywords"`
	ReferenceIDs []int    `json:"references"`
	Tags         []int    `json:"tags"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Record struct {
	ID        int       `json:"id"`
	PaperID   int       `json:"paper_id"`
	UserID    int       `json:"user_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Journal struct {
	ID    int         `json:"id"`
	Title string      `json:"title"`
	Link  string      `json:"link"`
	DBs   []JournalDB `json:"db"`
}

type Author struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	EnglishName   string     `json:"english_name"`
	Email         string     `json:"email"`
	ORCID         string     `json:"orcid"`
	AffiliationID int        `json:"affiliation"`
	Role          AuthorRole `json:"role"`
}

type Affiliation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type PaperTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AuthorRole string

func (_ AuthorRole) First() string                 { return "First author" }
func (_ AuthorRole) Corresponding() string         { return "Corresponding author" }
func (_ AuthorRole) Corperation() string           { return "Co-author" }
func (_ AuthorRole) PrincipalInvestigator() string { return "Principal investigator" }
func (_ AuthorRole) CoFirst() string               { return "Co-first author" }
func (_ AuthorRole) Other() string                 { return "Other author" }

type JournalDB string

func (_ JournalDB) KCI() string    { return "KCI" }
func (_ JournalDB) SCI() string    { return "SCI" }
func (_ JournalDB) SCIE() string   { return "SCIE Korea" }
func (_ JournalDB) SCOPUS() string { return "SCOPUS" }
func (_ JournalDB) IEEE() string   { return "IEEE" }
