package app

type Journal struct {
	ID    ID          `json:"id"`
	Title string      `json:"title"`
	Link  string      `json:"link"`
	DBs   []JournalDB `json:"db"`
}

type JournalDB string

const (
	KCI     JournalDB = "KCI"
	SCI     JournalDB = "SCI"
	SCIE    JournalDB = "SCIE"
	SCOPUS  JournalDB = "SCOPUS"
	IEEE    JournalDB = "IEEE"
	Unknown JournalDB = "Unknown"
	Other   JournalDB = "Other"
)
