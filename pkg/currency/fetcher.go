package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const NBU_API_URL = "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json"

// Rate represents the exchange rate data from NBU API
type Rate struct {
	R030         int     `json:"r030"`
	Text         string  `json:"txt"`
	Rate         float64 `json:"rate"`
	CurrencyCode string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}

// FetchCurrentRateNBU fetches the current exchange rate from USD to UAH from NBU API
func FetchCurrentRateNBU() (float64, error) {
	body, err := FetchData(NBU_API_URL)
	if err != nil {
		return 0, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(body)

	return FindUSDRateNBU(body)
}

// FetchData performs the HTTP GET request and returns the response body
func FetchData(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching data:", err)
		return nil, err
	}
	return resp.Body, nil
}

// FindUSDRateNBU decodes the JSON from an io.Reader and extracts the USD exchange rate from NBU data
func FindUSDRateNBU(data io.Reader) (float64, error) {
	var rates []Rate
	if err := json.NewDecoder(data).Decode(&rates); err != nil {
		return 0, fmt.Errorf("error decoding response: %v", err)
	}

	for _, rate := range rates {
		if rate.CurrencyCode == "USD" {
			return rate.Rate, nil
		}
	}

	return 0, fmt.Errorf("USD rate not found")
}
