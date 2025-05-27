package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pricetula/gaze-news-api/internal/api"
	"github.com/pricetula/gaze-news-api/internal/db/sqlxdb"
	"github.com/pricetula/gaze-news-api/internal/infrastructure/news"
	"github.com/pricetula/gaze-news-api/internal/infrastructure/uow"
	"github.com/pricetula/gaze-news-api/internal/utils"
)

func main() {
	// Create a new context
	ctx := context.Background()

	// Load configuration from various sources like environment variables or configuration files
	cfg, err := utils.NewConfig()
	if err != nil {
		panic(err)
	}

	// Setup the database connection
	db, err := sqlxdb.SetupDB(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// Setup UoW
	unitOfWork := uow.New(db)

	// Setup the news API client
	newsAPI := news.NewNews(cfg)

	// Setup fiber app
	app := fiber.New()

	// Group routes under /v1
	router := app.Group("/v1")

	// Inject DB into routes
	api.SetupRoutes(ctx, router, unitOfWork, newsAPI)

	// Start background fetch scheduler
	// go scheduler.Start(database)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err = app.Listen(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
