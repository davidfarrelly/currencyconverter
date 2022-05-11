package parser

import (
	"currency-converter/src/converter"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

/*
	Generates a conversion object from CLI input
*/
func ParseCliInput(base, target, date string, amount float64) *converter.Conversion {
	conversion := converter.Conversion{
		Base:   base,
		Target: target,
		Date:   date,
		Amount: amount}

	return &conversion
}

/*
	Generates a conversion object from file input
*/
func ParseFileInput(file string) (*converter.Conversion, error) {
	ext := filepath.Ext(file)
	var conversion *converter.Conversion

	convFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
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

	return conversion, nil
}
