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

	var products []*shoes.Product
	for _, url := range urls {
		variants, err := a.scraper.ScrapeBuyShoesVariants(
			ctx,
			url,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &shoes.Product{
			URL:      url,
			Variants: variants,
		})
	}
	return products, nil
}
