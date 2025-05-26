package uow

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pricetula/gaze-news-api/internal/domain"
	"github.com/pricetula/gaze-news-api/internal/repository"
)

type Repositories struct {
	ArticleRepository domain.ArticleRepository
	SourcesRepository domain.SourcesRepository
}

type UnitOfWork interface {
	Do(ctx context.Context, fn func(r *Repositories) error) error
}

type unitOfWork struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UnitOfWork {
	return &unitOfWork{db}
}

func (u *unitOfWork) Do(ctx context.Context, fn func(r *Repositories) error) error {
	// Start a new transaction
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	// Ensure the transaction is rolled back if the function returns an error
	defer tx.Rollback()

	// Create a new instance of Repositories with the transaction
	repos := &Repositories{
		ArticleRepository: repository.NewArticlesRepository(tx),
		SourcesRepository: repository.NewSourcesRepository(tx),
	}

	// Execute the function with the repositories. If it returns an error, rollback the transaction
	if err := fn(repos); err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction if everything went well
	return tx.Commit()
}
