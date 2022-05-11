package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RateInfo struct {
	Success   bool
	Timestamp string
	Base      string
	Date      string
	Rates     interface{}
}

const LATEST_URL = "/fixer/latest"
const HIST_URL = "/fixer/"

type Client interface {
	GetRate(base, target, date string) (RateInfo, error)
}

type ApiClient struct {
	BaseUrl string
}

func NewApiClient(baseUrl string) *ApiClient {
	return &ApiClient{
		BaseUrl: baseUrl,
	}
}

func (client ApiClient) GetRate(base, target, date string) (RateInfo, error) {
	var rateInfo RateInfo
	var url string

	if date == "" {
		url = client.BaseUrl + LATEST_URL
	} else {
		url = client.BaseUrl + HIST_URL + date
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("base", base)
	q.Add("symbols", target)

	req.URL.RawQuery = q.Encode()

	response, err := doRequest(req)
	if err != nil {
		return rateInfo, errors.New("error getting rate. " + err.Error())
	}

	json.Unmarshal(response, &rateInfo)

	return rateInfo, nil
}

func doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}

	apiKey, present := os.LookupEnv("API_KEY")
	if !present {
		log.Fatal("required ENV variable API_KEY not set.")
	}

	req.Header.Set("apikey", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error sending http request, status: " + resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil

}
