package test

import (
	"currency-converter/src/client"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Tests using a mock server
*/

func addRatesApiHandler(t *testing.T, mux *http.ServeMux, isError bool) {
	latestRates := make(map[string]interface{})
	latestRates["USD"] = 2.0
	latestRateInfo, _ := json.Marshal(client.RateInfo{Base: "EUR", Rates: latestRates})

	histRates := make(map[string]interface{})
	histRates["USD"] = 3.0
	historicalRateInfo, _ := json.Marshal(client.RateInfo{Base: "EUR", Rates: histRates, Date: "2000-01-01"})

	mux.HandleFunc("/fixer/latest", func(res http.ResponseWriter, req *http.Request) {
		if isError {
			res.WriteHeader(http.StatusBadRequest)
		} else {
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(latestRateInfo))
		}
	})

	mux.HandleFunc("/fixer/2000-01-01", func(res http.ResponseWriter, req *http.Request) {
		if isError {
			res.WriteHeader(http.StatusBadRequest)
		} else {
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(historicalRateInfo))
		}
	})

}

func TestGetLatestRateError(t *testing.T) {
	mux := http.NewServeMux()
	addRatesApiHandler(t, mux, true)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	client := client.NewApiClient(ts.URL)
	_, err := client.GetRate("EUR", "USD", "")

	assert.NotNil(t, err)
	assert.Equal(t, "error getting rate. error sending http request, status: 400 Bad Request", err.Error())
}

func TestGetLatestRate(t *testing.T) {
	mux := http.NewServeMux()
	addRatesApiHandler(t, mux, false)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	client := client.NewApiClient(ts.URL)
	latestRate, err := client.GetRate("EUR", "USD", "")
	assert.Nil(t, err)

	rateMap := latestRate.Rates.(map[string]interface{})
	rate := rateMap["USD"].(float64)

	assert.Equal(t, 2.0, rate)
}

func TestGetLatestHistoricalRate(t *testing.T) {
	mux := http.NewServeMux()
	addRatesApiHandler(t, mux, false)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	client := client.NewApiClient(ts.URL)
	latestRate, err := client.GetRate("EUR", "USD", "2000-01-01")
	assert.Nil(t, err)

	rateMap := latestRate.Rates.(map[string]interface{})
	rate := rateMap["USD"].(float64)

	assert.Equal(t, 3.0, rate)
	assert.Equal(t, "2000-01-01", latestRate.Date)
}

func TestGetLatestHistoriclaRateError(t *testing.T) {
	mux := http.NewServeMux()
	addRatesApiHandler(t, mux, true)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	client := client.NewApiClient(ts.URL)
	_, err := client.GetRate("EUR", "USD", "2000-01-01")

	assert.NotNil(t, err)
	assert.Equal(t, "error getting rate. error sending http request, status: 400 Bad Request", err.Error())
}
