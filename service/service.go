package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"stockpredictionscronjob/stringutil"
	"time"
)

var apiBaseAddres string = os.Getenv("STOCK_PREDICTIONS_API")
var apiKey string = os.Getenv("STOCK_PREDICTIONS_API_KEY")
var apiSecret string = os.Getenv("STOCK_PREDICTIONS_API_SECRET")
var pastStocksRef string = os.Getenv("PAST_STOCKS_REF")
var client = &http.Client{}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func doAuthenticatedRequest(method string, url string) (*http.Response, error) {

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(apiKey, apiSecret))
	resp, err := client.Do(req)

	return resp, err
}

func BootstrapFirstHistories() {

	stocks := GetNonCompliantPastDaysStocks()
	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/data​/stock-prices/history​/" + stocks[i])

		resp, err := doAuthenticatedRequest("POST", finalUrl)

		if err != nil {
			fmt.Println(err.Error())
		}
		if resp.Status != "200 OK" || err != nil {
			fmt.Println(stocks[i] + " could not been updated")
		}
		if resp.Status == "200 OK" {
			fmt.Println(stringutil.GetCurrentTimeStr() + stocks[i] + " has been bootstrapped sucessfully")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func SaveLastStockPrices() {
	stocks := GetSupportedStockPrices()

	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/data​/stock-prices​/" + stocks[i])

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Application has recoverd from error: ", r)
			}
		}()

		resp, err := doAuthenticatedRequest("POST", finalUrl)

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
	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/data​/stock-prices/" + stocks[i])

		resp, err := doAuthenticatedRequest("PUT", finalUrl)

		if resp.Status != "200 OK" || err != nil {
			fmt.Println(stringutil.GetCurrentTimeStr() + stocks[i] + " could not been updated")
		}
		if resp.Status == "200 OK" {
			fmt.Println(stringutil.GetCurrentTimeStr() + stocks[i] + " prices has been updated sucessfully")
		}
	}
}

func UpdatePrdictionLog() {
	stocks := GetSupportedStockPrices()
	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/stats/predictions/" + stocks[i])

		resp, err := doAuthenticatedRequest("PUT", finalUrl)

		if resp.Status != "200 OK" || err != nil {
			fmt.Println(stocks[i] + " could not been updated")
		}
		if resp.Status == "200 OK" {
			fmt.Println(stringutil.GetCurrentTimeStr() + stocks[i] + " log has been updated sucessfully")
		}
	}
}

func MakePredictions() {
	stocks := GetSupportedStockPrices()

	for i := 0; i < len(stocks); i++ {
		finalUrl := stringutil.CleanStr(apiBaseAddres + "​/prediction/nextday/" + stocks[i] + "?save_log=true")

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

func GetNonCompliantPastDaysStocks() []string {

	if len(pastStocksRef) == 0 {
		pastStocksRef = "40"
	}
	resp, err := http.Get(apiBaseAddres + "/data/supported-stocks/non-compliant/" + pastStocksRef)
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
