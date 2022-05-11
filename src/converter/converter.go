package converter

import (
	"currency-converter/src/client"
	"fmt"
)

type Converter struct {
	Client client.Client
}

type Conversion struct {
	Base   string
	Target string
	Amount float64
	Result float64
}

func NewConverter(client client.Client) Converter {
	return Converter{
		Client: client,
	}
}

func (converter *Converter) Convert(conversion *Conversion) error {
	rateInfo, err := converter.Client.GetRate(conversion.Base, conversion.Target)
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
