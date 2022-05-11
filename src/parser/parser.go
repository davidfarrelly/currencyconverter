package parser

import (
	"currency-converter/src/converter"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

/*
	Generates a conversion object from CLI input
*/
func ParseCliInput(base, target, date string, amount float64) (*converter.Conversion, error) {
	conversion := converter.Conversion{
		Base:   base,
		Target: target,
		Date:   date,
		Amount: amount}

	if err := checkInput(conversion); err != nil {
		return &conversion, errors.New("cli input error: " + err.Error())
	}

	return &conversion, nil
}

/*
	Generates a conversion object from file input
*/
func ParseFileInput(file string) (*converter.Conversion, error) {
	ext := filepath.Ext(file)
	var conversion *converter.Conversion

	convFile, err := os.Open(file)
	if err != nil {
		return conversion, errors.New("error reading input file: " + err.Error())
	}

	convBytes, _ := ioutil.ReadAll(convFile)

	switch ext {
	case ".json":
		if err := json.Unmarshal(convBytes, &conversion); err != nil {
			return conversion, errors.New("error unmarshalling json input file: " + err.Error())
		}
	case ".yaml":
		if err := yaml.Unmarshal(convBytes, &conversion); err != nil {
			return conversion, errors.New("error unmarshalling yaml input file: " + err.Error())
		}
	case ".yml":
		if err := yaml.Unmarshal(convBytes, &conversion); err != nil {
			return conversion, errors.New("error unmarshalling yaml input file: " + err.Error())
		}
	default:
		return conversion, errors.New(ext + " is not a supported file type.")
	}

	if err := checkInput(*conversion); err != nil {
		return conversion, errors.New("file input error: " + err.Error())
	}

	return conversion, nil
}

func checkInput(conversion converter.Conversion) error {
	if conversion.Base == "" {
		return errors.New("required base currency not supplied")
	} else if conversion.Target == "" {
		return errors.New("required target currency not supplied")
	} else if conversion.Amount == 0.0 {
		return errors.New("amount must be greater than 0.0")
	}

	return nil
}
