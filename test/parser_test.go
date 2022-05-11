package test

import (
	"currency-converter/src/parser"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCliInput(t *testing.T) {
	conversion, err := parser.ParseCliInput("EUR", "USD", "2000-01-01", 25.0)
	assert.Nil(t, err)

	assert.Equal(t, 25.0, conversion.Amount)
	assert.Equal(t, "EUR", conversion.Base)
	assert.Equal(t, "USD", conversion.Target)
	assert.Equal(t, "2000-01-01", conversion.Date)
}

func TestParseFileInputJson(t *testing.T) {
	file := "resources/conversion.json"
	conversion, err := parser.ParseFileInput(file)
	assert.Nil(t, err)

	assert.Equal(t, 50.0, conversion.Amount)
	assert.Equal(t, "EUR", conversion.Base)
	assert.Equal(t, "USD", conversion.Target)
	assert.Equal(t, "2000-01-01", conversion.Date)
}

func TestParseFileInputJsonInvalid(t *testing.T) {
	file := "resources/conversion-invalid.json"
	_, err := parser.ParseFileInput(file)
	assert.NotNil(t, err)

	correctErr := strings.Contains(err.Error(), "error unmarshalling json input file")

	assert.True(t, correctErr)
}

func TestParseFileInputYaml(t *testing.T) {
	file := "resources/conversion.yaml"
	conversion, err := parser.ParseFileInput(file)
	assert.Nil(t, err)

	assert.Equal(t, 25.0, conversion.Amount)
	assert.Equal(t, "GBP", conversion.Base)
	assert.Equal(t, "USD", conversion.Target)
	assert.Equal(t, "2000-01-01", conversion.Date)
}

func TestParseFileInputYamlInvalid(t *testing.T) {
	file := "resources/conversion-invalid.yaml"
	_, err := parser.ParseFileInput(file)
	assert.NotNil(t, err)

	correctErr := strings.Contains(err.Error(), "error unmarshalling yaml input file")

	assert.True(t, correctErr)
}

func TestParseFileInputUnssportedFileType(t *testing.T) {
	file := "resources/conversion.txt"
	_, err := parser.ParseFileInput(file)

	assert.NotNil(t, err)
	assert.Equal(t, ".txt is not a supported file type.", err.Error())
}

func TestParseFileInputFileNotExist(t *testing.T) {
	file := "resources/error.yaml"
	_, err := parser.ParseFileInput(file)

	assert.NotNil(t, err)
	assert.Equal(t, "error reading input file: open resources/error.yaml: The system cannot find the file specified.", err.Error())
}
