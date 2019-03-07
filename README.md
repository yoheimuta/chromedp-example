# chromedp-example

[![CircleCI](https://circleci.com/gh/yoheimuta/chromedp-example.svg?style=svg)](https://circleci.com/gh/yoheimuta/chromedp-example)

Scraping specific sneakers' bid price list in StockX:

- How to work:
  - Go to https://stockx.com/buy/air-jordan-1-retro-high-og-defiant-couture.
  - Wait to complete loading the confirmation page.
  - Click the confirm button.
  - Wait to complete loading the size list page.
  - Retrieve the sizes and prices.

## Setup

Set the project root path to GOPATH.

```bash
mkdir chromedp-example
cd chromedp-example
export GOPATH=$(pwd)
go get -d github.com/yoheimuta/chromedp-example # ignore a `no Go files` error.
cd src/github.com/yoheimuta/chromedp-example
```

Run your chrome headless-shell.

```bash
docker pull chromedp/headless-shell
docker run -d -p 9222:9222 --rm --name headless-shell chromedp/headless-shell
```

## Run

```bash
go run cmds/stockx/main.go
```

## Testing

```bash
go test -v -count 1 -timeout 240s -race ./...
```
