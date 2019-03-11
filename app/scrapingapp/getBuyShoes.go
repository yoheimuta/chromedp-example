package scrapingapp

import (
	"context"

	"github.com/yoheimuta/chromedp-example/domain/shoes"
)

func (a *App) GetBuyShoes(
	ctx context.Context,
) (
	[]*shoes.Product,
	error,
) {
	urls, err := a.db.GetShoesURLs(ctx)
	if err != nil {
		return nil, err
	}
	return a.scraper.ScrapeBuyShoesProducts(
		ctx,
		urls,
	)
}
