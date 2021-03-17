package main

import (
	"fmt"
	"stockpredictionscronjob/service"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Automation data handler is running...")
	c := cron.New(cron.WithSeconds())

	c.AddFunc("0 0 10 * * *", service.SaveLastStockPrices)
	c.AddFunc("0 * 11-18 * * *", service.UpdateLastStockPrices)

	c.AddFunc("0 20 18 * * *", service.UpdatePrdictionLog)
	c.AddFunc("0 30 18 * * *", service.MakePredictions)

	c.Start()

	fmt.Println(c.Entries())
	fmt.Scanln()
}
