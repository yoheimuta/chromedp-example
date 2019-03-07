package expchromedp

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

type Client struct {
	cdp *chromedp.CDP
}

type NewParams struct {
	log func(string, ...interface{})
}

type Option func(*NewParams)

func WithLog(log func(string, ...interface{})) Option {
	return func(params *NewParams) {
		params.log = log
	}
}

func NewClient(
	ctx context.Context,
	opts ...Option,
) (
	*Client,
	error,
) {
	params := &NewParams{}
	for _, opt := range opts {
		opt(params)
	}

	chromdpOpts := []chromedp.Option{
		chromedp.WithTargets(client.New().WatchPageTargets(ctx)),
	}
	if params.log != nil {
		chromdpOpts = append(chromdpOpts, chromedp.WithLog(params.log))
	}

	cdp, err := chromedp.New(
		ctx,
		chromdpOpts...,
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
