package app

import (
	"github.com/elecbug/jlum/internal/color"
)

type Paper struct {
	ID           ID               `json:"id"`
	Authors      []AuthorWithRole `json:"authors"`
	Title        string           `json:"title"`
	JournalID    ID               `json:"journal"`
	Volume       int              `json:"volume"`
	Issue        int              `json:"issue"`
	Page         Page             `json:"page"`
	Date         Date             `json:"date"`
	DOI          string           `json:"doi"`
	ISBN         string           `json:"isbn"`
	ISSN         string           `json:"issn"`
	EISSN        string           `json:"eissn"`
	Abstract     string           `json:"abstract"`
	FileData     []byte           `json:"file_data"`
	FileType     string           `json:"file_type"`
	Keywords     []string         `json:"keywords"`
	ReferenceIDs []ID             `json:"references"`
	TagIDs       []ID             `json:"tags"`
}

func (p *Paper) CiteIEEE() string {
}

type Page struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type PaperTag struct {
	ID    ID          `json:"id"`
	Name  string      `json:"name"`
	Color color.Color `json:"color"`
}
