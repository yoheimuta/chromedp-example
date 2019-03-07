package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/yoheimuta/chromedp-example/app/scrapingapp"
	"github.com/yoheimuta/chromedp-example/infra/expjson"
	"github.com/yoheimuta/chromedp-example/repository/db/expmockdb"
	"github.com/yoheimuta/chromedp-example/repository/scraper/expchromedp"
)

var (
	timeout = flag.Duration("timeout", 20*time.Second, "timeout")
)

func do() error {
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	scraper, err := expchromedp.NewClient(
		ctx,
		expchromedp.WithLog(log.Printf),
	)
	if err != nil {
		return err
	}
	defer func() {
		ctx2, cancel := context.WithTimeout(context.Background(), *timeout)
		defer cancel()

		err := scraper.Close(ctx2)
		if err != nil {
			log.Printf("Close err=%v\n", err)
		}
	}()
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
