package test

import (
	"currency-converter/src/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCliInput(t *testing.T) {
	conversion := parser.ParseCliInput("EUR", "USD", 25.0)

	assert.Equal(t, 25.0, conversion.Amount)
	assert.Equal(t, "EUR", conversion.Base)
	assert.Equal(t, "USD", conversion.Target)
}

func TestParseFileInputJson(t *testing.T) {
	file := "resources/conversion.json"
	conversion := parser.ParseFileInput(file)

	assert.Equal(t, 50.0, conversion.Amount)
	assert.Equal(t, "EUR", conversion.Base)
	assert.Equal(t, "USD", conversion.Target)
}
