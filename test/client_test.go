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

	mux.HandleFunc("/fixer/latest", func(res http.ResponseWriter, req *http.Request) {
		if isError {
			res.WriteHeader(http.StatusBadRequest)
		} else {
			res.WriteHeader(http.StatusOK)
			res.Write([]byte(latestRateInfo))
		}
	})

}

func TestGetLatestRate(t *testing.T) {
	mux := http.NewServeMux()
	addRatesApiHandler(t, mux, false)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	client := client.NewApiClient(ts.URL)
	latestRate, err := client.GetRate("EUR", "USD")
	assert.Nil(t, err)

	rateMap := latestRate.Rates.(map[string]interface{})
	rate := rateMap["USD"].(float64)

	assert.Equal(t, 2.0, rate)
}
