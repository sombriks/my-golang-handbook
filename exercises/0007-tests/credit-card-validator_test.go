package main

import "testing"

func TestInvalidBrand(t *testing.T) {

	var card CreditCardInfo = CreditCardInfo{
		"4417 1234 5678 9113",
		"Visa_",
	}

	_, err := card.IsValid()

	if err == nil {
		t.Fatal("The brand isn't supposed to exists")
	}

}

func TestInvalidBIN(t *testing.T) {

	card := CreditCardInfo{
		"3417 1234 5678 9113",
		Visa,
	}

	_, err := card.IsValid()

	if err == nil {
		t.Fatal("The BIN isn't correct")
	}

}

func TestInvalidMod10(t *testing.T) {

	var card CreditCardInfo

	card = CreditCardInfo{
		"4427 1234 5678 9113",
		Visa,
	}

	pass, _ := card.IsValid()

	if pass {
		t.Fatal("The Hans Peter Luhn algorithm isn't supposed to pass")
	}

}
