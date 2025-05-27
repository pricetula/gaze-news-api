package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pricetula/gaze-news-api/internal/domain"
)

type ArticlesRepository struct {
	db *sqlx.Tx
}

func NewArticlesRepository(db *sqlx.Tx) *ArticlesRepository {
	return &ArticlesRepository{db: db}
}

func (r *ArticlesRepository) GetArticlesByIDs(ids []uuid.UUID) ([]domain.Article, error) {
	// Build query which selects articles by their IDs and inner joins user_settings.
	query, args, err := sqlx.In(
		`SELECT
			article.id, article.title, article.description, article.url, article.url_to_image, article.published_at, article.content,
			author.id, author.name,
			source.id, source.name
		FROM articles AS article
		INNER JOIN authors AS author ON author.id = author_id
		INNER JOIN sources AS source ON source.id = source_id
		WHERE article.id IN (?)`,
		ids,
	)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	var articles []domain.Article

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		a := domain.Article{
			Source: &domain.Source{},
			Author: &domain.Author{},
		}
		err = rows.Scan(
			&a.ID,
			&a.Title,
			&a.Description,
			&a.URL,
			&a.URLToImage,
			&a.PublishedAt,
			&a.Content,
			&a.Author.ID,
			&a.Author.Name,
			&a.Source.ID,
			&a.Source.Name,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	return articles, nil
}

func (r *ArticlesRepository) GetArticles() ([]domain.Article, error) {
	// Build query which selects all articles and inner joins user_settings.
	query := `
		SELECT
			article.id, article.title, article.description, article.url, article.url_to_image, article.published_at, article.content,
			author.id, author.name,
			source.id, source.name
		FROM articles AS article
		INNER JOIN authors AS author ON author.id = author_id
		INNER JOIN sources AS source ON source.id = source_id
	`

	var articles []domain.Article

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		a := domain.Article{
			Source: &domain.Source{},
			Author: &domain.Author{},
		}
		err = rows.Scan(
			&a.ID,
			&a.Title,
			&a.Description,
			&a.URL,
			&a.URLToImage,
			&a.PublishedAt,
			&a.Content,
			&a.Author.ID,
			&a.Author.Name,
			&a.Source.ID,
			&a.Source.Name,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}
