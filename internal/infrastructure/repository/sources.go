package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pricetula/gaze-news-api/internal/domain"
)

type SourcesRepository struct {
	db *sqlx.Tx
}

func NewSourcesRepository(db *sqlx.Tx) *SourcesRepository {
	return &SourcesRepository{db: db}
}

func (r *SourcesRepository) GetSources() ([]domain.Source, error) {
	// Build query which selects sources.
	query := `SELECT
		source.id, source.name, source.description, source.url,
		category.id, category.name,
		language.id, language.name,
		country.id, country.name
		FROM sources AS source
		INNER JOIN categories AS category ON category.id = source.category_id
		INNER JOIN languages AS language ON language.id = source.language_id
		INNER JOIN countries AS country ON country.id = source.country_id`

	var sources []domain.Source

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := domain.Source{
			Category: &domain.Category{},
			Language: &domain.Language{},
			Country:  &domain.Country{},
		}
		err = rows.Scan(
			&s.ID,
			&s.Name,
			&s.Description,
			&s.URL,
			&s.Category.ID,
			&s.Category.Name,
			&s.Language.ID,
			&s.Language.Name,
			&s.Country.ID,
			&s.Country.Name,
		)
		if err != nil {
			return nil, err
		}
		sources = append(sources, s)
	}

	return sources, nil
}

func (r *SourcesRepository) AddSources(sources []*domain.Source) error {
	// Build query which inserts sources.
	query := `INSERT INTO sources
		(id, name, description, url, category_id, language_id, country_id)
		VALUES (:id, :name, :description, :url, :category_id, :language_id, :country_id)`

	_, err := r.db.NamedExec(query, sources)
	if err != nil {
		return err
	}

	return nil
}
