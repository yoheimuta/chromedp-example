package scrapingapp_test

import (
	"context"
	"testing"

	"github.com/yoheimuta/chromedp-example/app/scrapingapp"
	"github.com/yoheimuta/chromedp-example/repository/db/expmockdb"
	"github.com/yoheimuta/chromedp-example/repository/scraper/expchromedp"
	"github.com/yoheimuta/chromedp-example/repository/scraper/expfastchromedp"
)

type mockableDBClient struct {
	*expmockdb.Client
	urls []string
}

func (c *mockableDBClient) GetShoesURLs(
	context.Context,
) (
	[]string,
	error,
) {
	return c.urls, nil
}

func TestApp_GetBuyShoes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	scraper, err := expchromedp.NewClient(ctx)
	if err != nil {
		t.Errorf("got err %v", err)
		return
	}
	defer func() {
		err = scraper.Close(ctx)
		if err != nil {
			t.Errorf("got err %v", err)
		}
	}()

	fastScraper, err := expfastchromedp.NewClient(ctx)
	if err != nil {
		t.Errorf("got err %v", err)
		return
	}
	defer func() {
		err = fastScraper.Close(ctx)
		if err != nil {
			t.Errorf("got err %v", err)
		}
	}()

	db := &mockableDBClient{}

	for _, test := range []struct {
		name             string
		inputURLs        []string
		inputScraper     scrapingapp.Scraper
		wantVariantCount []int
	}{
		{
			name: "Got 1 url",
			inputURLs: []string{
				`https://stockx.com/buy/air-jordan-1-retro-high-og-defiant-couture`,
			},
			inputScraper:     scraper,
			wantVariantCount: []int{17},
		},
		{
			name: "Got 2 urls",
			inputURLs: []string{
				`https://stockx.com/buy/air-jordan-1-retro-high-og-defiant-couture`,
				`https://stockx.com/buy/adidas-yeezy-boost-700-salt`,
			},
			inputScraper:     scraper,
			wantVariantCount: []int{17, 25},
		},
		{
			name: "Got 1 url with the cookie to skip the confirmation page",
			inputURLs: []string{
				`https://stockx.com/buy/air-jordan-1-retro-high-og-defiant-couture`,
			},
			inputScraper:     fastScraper,
			wantVariantCount: []int{17},
		},
		{
			name: "Got 2 urls with the cookie to skip the confirmation page",
			inputURLs: []string{
				`https://stockx.com/buy/air-jordan-1-retro-high-og-defiant-couture`,
				`https://stockx.com/buy/adidas-yeezy-boost-700-salt`,
			},
			inputScraper:     fastScraper,
			wantVariantCount: []int{17, 25},
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			db.urls = test.inputURLs

			app := scrapingapp.NewApp(test.inputScraper, db)
			got, err := app.GetBuyShoes(ctx)
			if err != nil {
				t.Errorf("got err %v", err)
				return
			}

			if len(got) != len(test.wantVariantCount) {
				t.Errorf("got %d, but want %d", len(got), len(test.wantVariantCount))
				return
			}
			for i := 0; i < len(test.wantVariantCount); i++ {
				if len(got[i].Variants) != test.wantVariantCount[i] {
					t.Errorf("got %d, but want %d", len(got[i].Variants), test.wantVariantCount[i])
					return
				}
			}
		})
	}
}
