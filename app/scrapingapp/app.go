package scrapingapp

import (
	"context"

	"github.com/yoheimuta/chromedp-example/domain/shoes"
)

type Scraper interface {
	ScrapeBuyShoesVariants(
		ctx context.Context,
		shoesURL string,
	) (
		[]*shoes.Variant,
		error,
	)
}

type DB interface {
	GetShoesURLs(
		ctx context.Context,
	) (
		[]string,
		error,
	)
}

type App struct {
	scraper Scraper
	db      DB
}

func NewApp(
	scraper Scraper,
	db DB,
) *App {
	return &App{
		scraper: scraper,
		db:      db,
	}
}
