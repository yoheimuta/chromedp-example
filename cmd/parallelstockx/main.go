package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/yoheimuta/chromedp-example/repository/scraper/expparchromedp"

	"github.com/yoheimuta/chromedp-example/app/scrapingapp"
	"github.com/yoheimuta/chromedp-example/infra/expjson"
	"github.com/yoheimuta/chromedp-example/repository/db/expmockdb"
)

var (
	timeout = flag.Duration("timeout", 20*time.Second, "timeout")
)

func do() error {
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	scraper := expparchromedp.NewClient(
		expparchromedp.WithLog(log.Printf),
		expparchromedp.WithPortRange(10000, 20000),
	)
	db := expmockdb.NewClient()

	app := scrapingapp.NewApp(
		scraper,
		db,
	)

	products, err := app.GetBuyShoes(
		ctx,
	)
	if err != nil {
		return err
	}
	log.Printf("%v\n", expjson.PrettyFormat(products))
	return nil
}

func main() {
	err := do()
	if err != nil {
		log.Printf("err=%v\n", err)
	}
}
