package converter

import (
	"currency-converter/src/client"
	"fmt"
	"log"
)

type Converter struct {
	Client client.Client
}

type Conversion struct {
	Base   string
	Target string
	Amount float64
	Result float64
	Date   string
}

var supportedCurrencies = []string{"EUR", "USD", "GBP", "JPY", "AUD", "CHF", "CAD"}

func NewConverter(client client.Client) Converter {
	return Converter{
		Client: client,
	}
}

func (converter *Converter) Convert(conversion *Conversion) error {
	// Check if supported currency
	if !isSupportedCurrency(conversion.Base) {
		log.Fatal(conversion.Base + " is not a supported currency.")
	} else if !isSupportedCurrency(conversion.Target) {
		log.Fatal(conversion.Target + " is not a supported currency.")
	} else {
		fmt.Println("Converting from " + conversion.Base + " to " + conversion.Target)
	}

	rateInfo, err := converter.Client.GetRate(conversion.Base, conversion.Target, conversion.Date)
	if err != nil {
		return err
	}

	ratesMap := rateInfo.Rates.(map[string]interface{})
	rate := ratesMap[conversion.Target].(float64)

	conversion.Result = conversion.Amount * rate

	fmt.Println("Amount in " + conversion.Base + ": " + fmt.Sprintf("%f", conversion.Amount))
	fmt.Println("Amount in " + conversion.Target + ": " + fmt.Sprintf("%f", conversion.Result))

	return nil
}

func isSupportedCurrency(str string) bool {
	for _, v := range supportedCurrencies {
		if v == str {
			return true
		}
	}

	return false
}
