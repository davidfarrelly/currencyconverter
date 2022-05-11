package test

import (
	"currency-converter/src/client"
	"currency-converter/src/converter"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Tests using interface to mock api client (no need for mock server)
*/

type ApiClientMock struct{}

func (apiClientMock ApiClientMock) GetRate(base, target, date string) (client.RateInfo, error) {
	rates := make(map[string]interface{})

	if date == "" {
		// latest rate
		rates[target] = 2.0
	} else {
		// historical rate
		rates[target] = 3.0
	}

	return client.RateInfo{Base: base, Rates: rates}, nil
}

func TestConvert(t *testing.T) {
	apiClientMock := ApiClientMock{}

	conversion := converter.Conversion{
		Base:   "EUR",
		Target: "USD",
		Amount: 10,
		Date:   "",
	}

	converter := converter.NewConverter(apiClientMock)
	converter.Convert(&conversion)

	assert.Equal(t, 20.0, conversion.Result)
}

func TestConvertHistorical(t *testing.T) {
	apiClientMock := ApiClientMock{}

	conversion := converter.Conversion{
		Base:   "EUR",
		Target: "USD",
		Amount: 10,
		Date:   "2000-01-01",
	}

	converter := converter.NewConverter(apiClientMock)
	converter.Convert(&conversion)

	assert.Equal(t, 30.0, conversion.Result)
}
