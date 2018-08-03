package schema

import "time"

// TechArticle contains all data and metadata of a given page.
type TechArticle struct {
	JsonLd
	Name                    string          `json:"name"`
	Text                    string          `json:"text"`
	ArticleBody             string          `json:"articleBody"`
	About                   string          `json:"about,omitempty"`
	Author                  []Person        `json:"author,omitempty"`
	Sponsor                 []Person        `json:"sponsor,omitempty"`
	SubjectOf               []CreativeWork  `json:"subjectOf,omitempty"`
	Keywords                Text            `json:"keywords,omitempty"`
	Description             Text            `json:"description,omitempty"`
	SourceOrganization      Organization    `json:"sourceOrganization,omitempty"`
	LearningResourceType    Text            `json:"learningResourceType,omitempty"`
	IsPartOf                CreativeWork    `json:"`json:"learningResourceType,omitempty"`
	isAccessibleForFree     Boolean         `json:"`json:"isAccessibleForFree,omitempty"`
	ExampleOfWork           CreativeWork    `json:"`json:"exampleOfWork,omitempty"`
	Audience                CreativeWork    `json:"`json:"audience,omitempty"`
	Author                  []Person        `json:"`json:"author,omitempty"`
	creator                 Person          `json:"`json:"creator,omitempty"`
	DateCreated             *time.Time      `json:"dateCreated,omitempty"`
	DateModified            *time.Time      `json:"dateModified,omitempty"`
	InLanguage              string          `json:"inLanguage,omitempty"`
	License                 string          `json:"license,omitempty"`
	Publisher               string          `json:"publisher,omitempty"`
}

// NewTechArticle initializes a new TechArticle instance with some sensitive default values.
func NewTechArticle() *TechArticle {
	return &TechArticle{JsonLd: JsonLd{Context: "http://schema.org", Type: "TechArticle"}}
}
