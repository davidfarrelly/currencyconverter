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

	file := fileCommand.String("input", "", "Path to the JSON/YAML input file. (required)")

	var conversion *converter.Conversion
	var err error

	if len(os.Args) < 2 {
		log.Fatal("required subcommand not provided")
	}

	switch os.Args[1] {
	case "cli":
		cliCommand.Parse(os.Args[2:])
		conversion, err = parser.ParseCliInput(*baseCurrency, *targetCurrency, *date, *amount)
		if err != nil {
			log.Fatal("error parsing cli input: " + err.Error())
		}
	case "file":
		fileCommand.Parse(os.Args[2:])
		conversion, err = parser.ParseFileInput(*file)
		if err != nil {
			log.Fatal("error parsing file input: " + err.Error())
		}
	default:
		log.Fatal("unknown subcommand provided, see -h for more details")
	}

	client := client.NewApiClient(BASE_URL)
	converter := converter.NewConverter(*client)

	if err := converter.Convert(conversion); err != nil {
		log.Fatal("error converting currency: " + err.Error())
	}
}
