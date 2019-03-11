package expparchromedp

import (
	"context"
	"fmt"

	"github.com/chromedp/cdproto/cdp"
	"github.com/yoheimuta/chromedp-example/infra/expchromedp"

	"github.com/chromedp/chromedp"
	"github.com/yoheimuta/chromedp-example/domain/shoes"
)

func (c *Client) ScrapeBuyShoesProducts(
	ctx context.Context,
	shoesURLs []string,
) (
	[]*shoes.Product,
	error,
) {
	// create pool
	pool, err := chromedp.NewPool(c.poolOptions()...)
	if err != nil {
		return nil, err
	}
	// shutdown pool
	defer func() {
		serr := pool.Shutdown()
		if err == nil && serr != nil {
			err = serr
		}
	}()

	// loop over the URLs
	productChan := make(chan *shoes.Product)
	errorChan := make(chan error)
	for _, url := range shoesURLs {
		url := url
		go func() {
			vs, err := c.scrapeBuyShoesVariants(
				ctx,
				pool,
				url,
			)
			if err != nil {
				errorChan <- err
			}
			productChan <- &shoes.Product{
				URL:      url,
				Variants: vs,
			}
		}()
	}

	// wait for to finish
	var products []*shoes.Product
	for i := 0; i < len(shoesURLs); i++ {
		select {
		case p := <-productChan:
			products = append(products, p)
		case err := <-errorChan:
			return nil, err
		}
	}
	close(productChan)
	close(errorChan)

	return products, nil
}

func (c *Client) scrapeBuyShoesVariants(
	ctx context.Context,
	pool *chromedp.Pool,
	shoesURL string,
) (
	_ []*shoes.Variant,
	err error,
) {
	// allocate
	r, err := pool.Allocate(ctx)
	if err != nil {
		return nil, fmt.Errorf("url `%s` error: %v", shoesURL, err)
	}
	defer func() {
		_ = r.Release()
	}()

	// run tasks
	var sizes []*cdp.Node
	var prices []*cdp.Node
	confirmSel := `//*[@id="bottom-bar-root"]/div/div/button[2]`
	sizesSel := `//div[@class='tile-inner']/div[@class='tile-value']`
	sizeTextsSel := sizesSel + `/text()`
	priceTextsSel := `//div[@class='tile-inner']/div[@class='tile-subvalue']/div/text()`
	err = r.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(shoesURL),
		chromedp.WaitVisible(confirmSel),
		chromedp.Click(confirmSel),
		chromedp.WaitVisible(sizesSel),
		chromedp.Nodes(sizeTextsSel, &sizes),
		chromedp.Nodes(priceTextsSel, &prices),
	})
	if err != nil {
		return nil, err
	}

	return shoes.NewVariants(
		expchromedp.NodeValues(sizes),
		expchromedp.NodeValues(prices),
	)
}
