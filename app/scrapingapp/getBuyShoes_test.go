package scrapingapp_test

import (
	"context"
	"log"
	"testing"

	"github.com/yoheimuta/chromedp-example/repository/scraper/expcncchromedp"

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
	scraper, err := expchromedp.NewClient(ctx, expchromedp.WithLog(log.Printf))
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

	fastScraper, err := expfastchromedp.NewClient(ctx, expfastchromedp.WithLog(log.Printf))
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

	concurrentScraper := expcncchromedp.NewClient(
		expcncchromedp.WithLog(log.Printf),
		expcncchromedp.WithPortRange(10000, 20000),
	)

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
		{
			name: "Got 1 url with concurrent mode",
			inputURLs: []string{
				`https://stockx.com/buy/air-jordan-1-retro-high-og-defiant-couture`,
			},
			inputScraper:     concurrentScraper,
			wantVariantCount: []int{17},
		},
		{
			name: "Got 2 urls with concurrent mode",
			inputURLs: []string{
				`https://stockx.com/buy/air-jordan-1-retro-high-og-defiant-couture`,
				`https://stockx.com/buy/adidas-yeezy-boost-700-salt`,
			},
			inputScraper:     concurrentScraper,
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

			gotM := make(map[string]int)
			for _, p := range got {
				gotM[p.URL] = len(p.Variants)
			}
			for i, url := range test.inputURLs {
				l, ok := gotM[url]
				if !ok {
					t.Errorf("not found %s in map", url)
					continue
				}
				if l != test.wantVariantCount[i] {
					t.Errorf("got %d, but want %d", l, test.wantVariantCount[i])
				}
			}
		})
	}
}
