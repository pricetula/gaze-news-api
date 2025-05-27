package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pricetula/gaze-news-api/internal/domain"
	"github.com/pricetula/gaze-news-api/internal/infrastructure/uow"
	"github.com/pricetula/gaze-news-api/internal/utils"
)

func getArticles(ctx context.Context, unitOfWork uow.UnitOfWork) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Extract the ids parameter from the query string
		ids := c.Query("ids")
		articles := []domain.Article{}

		// Check if the ids parameter is empty
		if ids == "" {
			if err := unitOfWork.Do(ctx, func(r *uow.Repositories) error {
				// Get articles from the database by their IDs
				f, err := r.ArticleRepository.GetArticles()
				if err != nil {
					return err
				}
				articles = f
				return nil
			}); err != nil {
				// Handle any errors that occurred during the database operation
				return fiber.NewError(fiber.StatusExpectationFailed, err.Error())
			}
			// Return the articles as JSON
			return c.JSON(articles)
		}

		// Define a slice of UUIDs to hold the parsed IDs
		var articleIDs []uuid.UUID

		// Split the ids string into a slice of strings
		for _, id := range utils.SpltStr2Slc(ids, ",") {
			// Parse each ID into a UUID
			articleUUID, err := uuid.Parse(id)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid UUID format")
			}

			// Append the parsed UUID to the slice
			articleIDs = append(articleIDs, articleUUID)
		}

		if err := unitOfWork.Do(ctx, func(r *uow.Repositories) error {
			// Get articles from the database by their IDs
			f, err := r.ArticleRepository.GetArticlesByIDs(articleIDs)
			if err != nil {
				return err
			}
			articles = f
			return nil
		}); err != nil {
			// Handle any errors that occurred during the database operation
			return fiber.NewError(fiber.StatusExpectationFailed, err.Error())
		}
		// Return the articles as JSON
		return c.JSON(articles)
	}
}
