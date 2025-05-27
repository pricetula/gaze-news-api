package news

import (
	"net/http"
	"time"

	"github.com/pricetula/gaze-news-api/internal/utils"
)

const (
	baseURL = "https://newsapi.org/v2"
)

type ArticleSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	Source      ArticleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	URLToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type ArticlesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

// SourcesResponse represents the response from the NewsAPI sources endpoint.
type SourcesResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

// News represents the NewsAPI client.
type News struct {
	apiKey     string
	httpClient *http.Client // Use a shared HTTP client for efficiency and proper timeout handling
}

// New creates a new NewsAPI client.
// It takes a Config struct to get the API key and initializes an HTTP client.
func NewNews(cfg *utils.Config) *News {
	return &News{
		apiKey: cfg.NEWS_API_KEY,
		httpClient: &http.Client{
			Timeout: 10 * time.Second, // Set a reasonable timeout for API calls
		},
	}
}
