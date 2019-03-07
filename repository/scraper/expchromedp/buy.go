package expchromedp

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
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
	var sizes []*cdp.Node
	var prices []*cdp.Node
	confirmSel := `//*[@id="bottom-bar-root"]/div/div/button[2]`
	sizesSel := `//div[@class='tile-inner']/div[@class='tile-value']`
	sizeTextsSel := sizesSel + `/text()`
	priceTextsSel := `//div[@class='tile-inner']/div[@class='tile-subvalue']/div/text()`
	err := c.cdp.Run(ctx, chromedp.Tasks{
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
