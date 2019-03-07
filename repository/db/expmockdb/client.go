package expmockdb

import "context"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetShoesURLs(
	context.Context,
) (
	[]string,
	error,
) {
	return []string{
		`https://stockx.com/buy/adidas-yeezy-boost-700-salt`,
	}, nil
}
