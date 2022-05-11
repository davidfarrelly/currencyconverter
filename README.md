# GoLang Currency Converter

This project contains a currency converter built in GoLang. Using the Fixer API (https://fixer.io/) to get the latest and historical rates. 

Generally would get clarification on requirements from the Product Owner but for this challenge, some assumptions were made:

 - All input options CLI, JSON and YAML should be supported
 - Storing the API Key as an environment variable would be sufficient
 - Only a single target currency would be supplied (Fixer API supports multiple)

# Release

- [v1.0.0](https://github.com/davidfarrelly/currencyconverter/releases/tag/v1.0.0)

# Building

The following command can be used to build an executable for Windows:

```bash
set GOOS=windows
set GOARCH=amd64

cd src/
go build -o converter.exe
```

Or on Linux:

```bash
export GOOS=linux
export GOARCH=amd64

cd src/
go build -o converter
```

# Usage

The application supports two different input options, CLI and file based input can be used.

Currently supported currencies: ```["EUR", "USD", "GBP", "JPY", "AUD", "CHF", "CAD"]```

## Pre-requisites

- ```API_KEY``` environment variable must be set. See [Fixer Docs](https://fixer.io/documentation) for details on getting an API Key. [ISSUE-1](https://github.com/davidfarrelly/currencyconverter/issues/1)


## CLI Input

```
Usage of cli:
  -amount float
        The amount to be converted. (required)
  -base string
        Base currency to be converted from {EUR|USD|GBP|JPY|AUD|CHF|CAD}. (required)
  -date string
        The historical date to be used for conversion. {YYYY-MM-DD}
  -target string
        Target currency to be converted to {EUR|USD|GBP|JPY|AUD|CHF|CAD}. (required)
```

e.g

```bash
./converter.exe cli -amount 50.0 -base EUR -target USD -date 2000-01-01
```

## File Input

```
Usage of file:
  -input string
        Path to the JSON/YAML input file. (required)
```

e.g

```bash
./converter.exe file -input example/conversion.json
```

### File Input Format

Parameters can be supplied as either a JSON or YAML/YML file. Must be off the following format:

#### **`conversion.json`**
```json
{
    "base": "EUR",
    "target": "USD",
    "amount": 50.0,
    "date": "2000-01-01"
}
```

#### **`conversion.yaml`**
```yaml
base: "EUR"
target: "USD"
amount: 50.0
date: "2000-01-01"
```

# Testing

Unit tests can be run wih:

```bash
cd test/
go test -v
```

# Potential Improvements

 - Could be extended to allow the user to input multiple target currencies to be converted to.
 - Required API Key could be stored in a more secure way, an environment variable was used for this project for simplicity, as it is more secure than embedding the key within the codebase.
 - Configuration could be improved, config such as Fixer URL and supported currencies are hardcoded within the application. This could all be read in from a config file instead.
 - Application currently handles a single conversion then exits, could develop it further to continue running and handle consecutive conversions. 
