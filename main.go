package main

import (
	"fmt"
	"stockpredictionscronjob/service"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Automation data handler is running...")

	sptz, _ := time.LoadLocation("America/Sao_Paulo")
	c := cron.New(cron.WithLocation(sptz))

	service.BootstrapFirstHistories()
	fmt.Println("Bootstrap finished")

	c.AddFunc("15 10 * * MON-FRI", service.SaveLastStockPrices)

	c.AddFunc("0/2 11-17 * * MON-FRI", service.UpdateLastStockPrices)
	c.AddFunc("10 18 * * MON-FRI", service.UpdateLastStockPrices)

	c.AddFunc("15 18 * * MON-FRI", service.MakePredictions)
	c.AddFunc("30 18 * * MON-FRI", service.UpdatePrdictionLog)

	fmt.Println(c.Entries())

	c.Run()
}
