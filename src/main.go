package main

import (
	"currency-converter/src/client"
	"currency-converter/src/converter"
	"currency-converter/src/parser"
	"flag"
	"log"
	"os"
)

const BASE_URL = "https://api.apilayer.com"

func main() {

	// Declare sub-commands
	cliCommand := flag.NewFlagSet("cli", flag.ExitOnError)
	fileCommand := flag.NewFlagSet("file", flag.ExitOnError)

	// Declare command-line flags
	baseCurrency := cliCommand.String("base", "", "Base currency to be converted from {EUR|USD|GBP|JPY|AUD|CHF|CAD}. (required)")
	targetCurrency := cliCommand.String("target", "", "Target currency to be converted to {EUR|USD|GBP|JPY|AUD|CHF|CAD}. (required)")
	amount := cliCommand.Float64("amount", 0.0, "The amount to be converted. (required)")
	date := cliCommand.String("date", "", "The historical date to be used for conversion {YYYY-MM-DD}.")

	file := fileCommand.String("input", "", "JSON/YAML file input. (required)")

	var conversion *converter.Conversion

	switch os.Args[1] {
	case "cli":
		cliCommand.Parse(os.Args[2:])
		conversion = parser.ParseCliInput(*baseCurrency, *targetCurrency, *date, *amount)
	case "file":
		fileCommand.Parse(os.Args[2:])
		conversion = parser.ParseFileInput(*file)
	default:
		log.Fatal("required subcommand not supplied. See -h for usage.")
	}

	client := client.NewApiClient(BASE_URL)
	converter := converter.NewConverter(*client)

	if err := converter.Convert(conversion); err != nil {
		log.Fatal("error converting currency: " + err.Error())
	}
}
