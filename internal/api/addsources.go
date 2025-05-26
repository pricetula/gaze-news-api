package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pricetula/gaze-news-api/internal/domain"
	"github.com/pricetula/gaze-news-api/internal/news"
	"github.com/pricetula/gaze-news-api/internal/uow"
)

func addsources(ctx context.Context, unitOfWork uow.UnitOfWork, newsAPI *news.News) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get sources from the NewsAPI
		s, err := newsAPI.GetSources()
		if err != nil {
			// Handle the error appropriately
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		// Append each source to the domain.Source slice
		l := []*domain.Source{}
		for _, source := range s {
			l = append(l, &domain.Source{
				ID:          source.ID,
				Name:        source.Name,
				Description: source.Description,
				URL:         source.URL,
				CategoryID:  source.Category,
				LanguageID:  source.Language,
				CountryID:   source.Country,
			})
		}

		// Use the unitOfWork to bulk insert the sources into the database
		if err := unitOfWork.Do(ctx, func(r *uow.Repositories) error {
			// Store the sources in the database
			err := r.SourcesRepository.AddSources(l)
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			// Handle any errors that occurred during the database operation
			return fiber.NewError(fiber.StatusExpectationFailed, err.Error())
		}
		// Return the articles as JSON
		return c.JSON(l)
	}
}
