package helpers

import (
	"testing"
)

func TestMasDe100(t *testing.T) {
	resp := "Hay más de 100 resultados"
	number, isLimit := ParseResponse(resp)
	if number != 100 && !isLimit {
		t.Error("Expected: ", 100, "isLimit:", true)
	}
}

func TestInvoiceNumber(t *testing.T) {
	resp := "80"
	number, isLimit := ParseResponse(resp)

	if number != 80 && isLimit {
		t.Error("Expected:", 80, "isLimit:", false)
	}
}
