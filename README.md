# chromedp-example

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
