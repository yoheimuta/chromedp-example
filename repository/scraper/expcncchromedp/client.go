package expcncchromedp

import (
	"github.com/chromedp/chromedp"
	"github.com/yoheimuta/chromedp-example/infra/expmath"
)

type Client struct {
	log       func(string, ...interface{})
	portRange *expmath.Range
}

type Option func(*Client)

func WithLog(log func(string, ...interface{})) Option {
	return func(client *Client) {
		client.log = log
	}
}

func WithPortRange(start, end int) Option {
	return func(client *Client) {
		client.portRange = &expmath.Range{
			Start: start,
			End:   end,
		}
	}
}

func NewClient(opts ...Option) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) poolOptions() []chromedp.PoolOption {
	var opts []chromedp.PoolOption
	if c.log != nil {
		opts = append(opts, chromedp.PoolLog(c.log, c.log, c.log))
	}
	if c.portRange != nil {
		opts = append(opts, chromedp.PortRange(c.portRange.Start, c.portRange.End))
	}
	return opts
}
