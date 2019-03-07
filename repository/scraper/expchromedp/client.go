package expchromedp

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

type Client struct {
	cdp *chromedp.CDP
}

func NewClient(
	ctx context.Context,
) (
	*Client,
	error,
) {
	cdp, err := chromedp.New(
		ctx,
		chromedp.WithTargets(client.New().WatchPageTargets(ctx)),
		chromedp.WithLog(log.Printf),
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		cdp: cdp,
	}, nil
}

func (c *Client) Close(
	ctx context.Context,
) error {
	// shutdown chrome
	err := c.cdp.Shutdown(ctx)
	if err != nil {
		return err
	}

	// wait for chrome to finish
	err = c.cdp.Wait()
	if err != nil {
		return err
	}
	return nil
}
