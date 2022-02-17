package main

import (
	"fmt"
	"os"
	"stockpredictionscronjob/service"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/robfig/cron/v3"
)

func main() {
	var newRelicLicenseKey string = os.Getenv("NEW_RELIC_LICENSE_KEY")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("stock-predictions-cronjob"),
		newrelic.ConfigLicense(newRelicLicenseKey),
	)

	service.SetNewRelicAgent(app)

	if err == nil {
		fmt.Println("New Relic agent could not been loaded")
	}

	fmt.Println("Automation data handler is running...")

	sptz, _ := time.LoadLocation("America/Sao_Paulo")
	c := cron.New(cron.WithLocation(sptz))

	service.BootstrapFirstHistories()
	fmt.Println("Bootstrap finished")

	c.AddFunc("15 10 * * MON-FRI", service.SaveLastStockPrices)

	c.AddFunc("0/5 11-17 * * MON-FRI", service.UpdateLastStockPrices)
	c.AddFunc("10 18 * * MON-FRI", service.UpdateLastStockPrices)

	c.AddFunc("15 18 * * MON-FRI", service.MakePredictions)
	c.AddFunc("30 18 * * MON-FRI", service.UpdatePrdictionLog)

	fmt.Println(c.Entries())

	c.Run()
}
