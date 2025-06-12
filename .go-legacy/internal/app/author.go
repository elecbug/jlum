package app

var AuthorMap map[ID]Author
var AffiliationMap map[ID]Affiliation

type Author struct {
	ID            ID     `json:"id"`
	Name          string `json:"name"`
	EnglishName   string `json:"english_name"`
	Email         string `json:"email"`
	ORCID         string `json:"orcid"`
	AffiliationID ID     `json:"affiliation"`
}

type AuthorWithRole struct {
	AuthorID   ID         `json:"author_id"`
	AuthorRole AuthorRole `json:"author_role"`
}

func (a Author) CiteIEEE() string {
	if a.EnglishName != "" {
		return a.EnglishName
	} else {
		return a.Name
	}
}

func (a AuthorWithRole) CiteIEEE() string {
	if author, ok := AuthorMap[a.AuthorID]; ok {
		return author.CiteIEEE()
	} else {
		return ""
	}
}

type Affiliation struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type AuthorRole string

const (
	FirstAuthor           AuthorRole = "First author"
	CoFirstAuthor         AuthorRole = "Co-first author"
	CoAuthor              AuthorRole = "Co-author"
	CorrespondingAuthor   AuthorRole = "Corresponding author"
	PrincipalInvestigator AuthorRole = "Principal investigator"
	OtherAuthor           AuthorRole = "Other author"
)
