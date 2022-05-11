package parser

import (
	"currency-converter/src/converter"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
func ParseFileInput(file string) *converter.Conversion {
	ext := filepath.Ext(file)
	var conversion *converter.Conversion

	convFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	convBytes, _ := ioutil.ReadAll(convFile)

	switch ext {
	case ".json":
		json.Unmarshal(convBytes, &conversion)
	}

	return conversion
}
