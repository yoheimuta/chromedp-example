package expfastchromedp

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/yoheimuta/chromedp-example/domain/shoes"
	"github.com/yoheimuta/chromedp-example/infra/expchromedp"
)

func (c *Client) ScrapeBuyShoesVariants(
	ctx context.Context,
	shoesURL string,
) (
	[]*shoes.Variant,
	error,
) {
	u, err := url.Parse(shoesURL)
	if err != nil {
		return nil, err
	}

	var sizes []*cdp.Node
	var prices []*cdp.Node
	sizesSel := `//div[@class='tile-inner']/div[@class='tile-value']`
	sizeTextsSel := sizesSel + `/text()`
	priceTextsSel := `//div[@class='tile-inner']/div[@class='tile-subvalue']/div/text()`
	err = c.cdp.Run(ctx, chromedp.Tasks{
		chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
			success, err := network.SetCookie("stockx_seen_bid_new_info", "true").
				WithDomain(u.Hostname()).
				Do(ctxt, h)
			if err != nil {
				return err
			}
			if !success {
				return fmt.Errorf("could not set cookie")
			}
			return nil
		}),
		chromedp.Navigate(shoesURL),
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
