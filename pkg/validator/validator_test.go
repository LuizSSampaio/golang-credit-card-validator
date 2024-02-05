package validator

import (
	"testing"
)

func TestValidCard(t *testing.T) {
	cardNumber := "79927398713"
	result, err := Validate(cardNumber)
	if !result {
		t.Errorf("Validate(%s) returned false, expected true.", cardNumber)
	} else if err != nil {
		t.Errorf("An unexpected error ocurred. Error: %s.", err)
	}
}

func TestInvalidCard(t *testing.T) {
	cardNumber := "79927398711"
	result, err := Validate(cardNumber)
	if result {
		t.Errorf("Validate(%s) returned true, expected false.", cardNumber)
	} else if err != nil {
		t.Errorf("An unexpected error ocurred. Error: %s.", err)
	}
}

func TestInvalidCardNumber(t *testing.T) {
	cardNumber := "aaaaaa"
	_, err := Validate(cardNumber)
	if err == nil {
		t.Error("Missing expected invalid card number error")
	}
}
