package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pricetula/gaze-news-api/internal/infrastructure/news"
	"github.com/pricetula/gaze-news-api/internal/infrastructure/uow"
)

// SetupRoutes registers all API routes
func SetupRoutes(ctx context.Context, router fiber.Router, unitOfWork uow.UnitOfWork, newsAPI *news.News) {
	// Pass the DB to each handler
	router.Get("/articles", getArticles(ctx, unitOfWork))
	router.Post("/sources", addsources(ctx, unitOfWork, newsAPI))
}
