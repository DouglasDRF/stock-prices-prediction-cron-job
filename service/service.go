package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"stockpredictionscronjob/stringutil"
	"strings"
	"time"
)

var apiBaseAddres string = os.Getenv("STOCK_PREDICTIONS_API")

func BootstrapFirstHistories() {
	stocks := GetSupportedStockPrices()
	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/data​/fill-history​/" + stocks[i])

		resp, err := http.Post(finalUrl, "application/json", strings.NewReader(""))
		if err != nil {
			fmt.Println(err.Error())
		}
		if resp.Status != "200 OK" || err != nil {
			fmt.Println(stocks[i] + " could not been updated")
		}
		if resp.Status == "200 OK" {
			fmt.Println(stocks[i] + " has been updated sucessfully")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func SaveLastStockPrices() {
	stocks := GetSupportedStockPrices()

	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/data​/save-last​/" + stocks[i])

		resp, err := http.Post(finalUrl, "application/json", strings.NewReader(""))
		if resp.Status != "200 OK" || err != nil {
			fmt.Println(err.Error())
			fmt.Println(stocks[i] + " could not been saved")
		}
		if resp.Status == "200 OK" {
			fmt.Println(stocks[i] + " has been saved sucessfully")
		}
	}
}

func UpdateLastStockPrices() {
	stocks := GetSupportedStockPrices()
	client := &http.Client{}
	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/data​/update-last​/" + stocks[i])

		request, err := http.NewRequest("PUT", finalUrl, strings.NewReader(""))
		if err != nil {
			fmt.Println(err.Error())
		}
		resp, err := client.Do(request)

		if resp.Status != "200 OK" || err != nil {
			fmt.Println(stocks[i] + " could not been updated")
		}
		if resp.Status == "200 OK" {
			fmt.Println(stocks[i] + " prices has been updated sucessfully")
		}
	}
}

func UpdatePrdictionLog() {
	stocks := GetSupportedStockPrices()
	client := &http.Client{}
	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/stats/" + stocks[i])

		request, err := http.NewRequest("PUT", finalUrl, strings.NewReader(""))
		if err != nil {
			fmt.Println(err.Error())
		}
		resp, err := client.Do(request)

		if resp.Status != "200 OK" || err != nil {
			fmt.Println(stocks[i] + " could not been updated")
		}
		if resp.Status == "200 OK" {
			fmt.Println(stocks[i] + " log has been updated sucessfully")
		}
	}
}

func MakePredictions() {
	stocks := GetSupportedStockPrices()

	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/predict/nextday/" + stocks[i])

		resp, err := http.Get(finalUrl)
		if err != nil || resp.Status != "200 OK" {
			fmt.Println(err.Error())
		}
		if resp.Status == "200 OK" {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(bodyBytes))
		}
	}
}

func GetSupportedStockPrices() []string {

	resp, err := http.Get(apiBaseAddres + "/data/supported-stocks")
	var stocks []string

	if err != nil {
		fmt.Println(err.Error())
		return stocks
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &stocks)
	return stocks
}
