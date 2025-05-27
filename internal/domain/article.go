package domain

import "github.com/google/uuid"

type Article struct {
	ID          string  `db:"id" json:"id" validate:"uuid"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	URL         string  `db:"url" json:"url"`
	URLToImage  string  `db:"url_to_image" json:"url_to_image"`
	PublishedAt string  `db:"published_at" json:"published_at"`
	Content     string  `db:"content" json:"content"`
	Author      *Author `json:"author"`
	Source      *Source `json:"source"`
}

type ArticleRepository interface {
	GetArticlesByIDs(ids []uuid.UUID) ([]Article, error)
	GetArticles() ([]Article, error)
}
