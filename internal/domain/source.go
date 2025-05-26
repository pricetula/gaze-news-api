package domain

type Source struct {
	ID          string    `db:"id" json:"id" validate:"uuid"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	URL         string    `db:"url" json:"url"`
	CategoryID  string    `db:"category_id" json:"category_id"`
	LanguageID  string    `db:"language_id" json:"language_id"`
	CountryID   string    `db:"country_id" json:"country_id"`
	Category    *Category `db:"category" json:"category"`
	Language    *Language `db:"language" json:"language"`
	Country     *Country  `db:"country" json:"country"`
}

type SourcesRepository interface {
	GetSources() ([]Source, error)
	AddSources([]*Source) error
}
