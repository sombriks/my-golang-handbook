package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BrandNames string

const (
	AmericanExpress BrandNames = "American Express"
	Visa            BrandNames = "Visa"
)

var brands map[BrandNames]string = map[BrandNames]string{
	AmericanExpress: "^34.*|^37.*",
	Visa:            "^4",
}

type CreditCardInfo struct {
	number string
	brand  BrandNames
}

// Name starts with capital letter so the function is exported
func (c CreditCardInfo) IsValid() (bool, error) {

	// check brand
	brand, exists := brands[c.brand]
	if !exists {
		return false, errors.New("this is not a valid brand")

	}

	// check BIN
	reg, _ := regexp.Compile(brand)
	if !reg.MatchString(c.number) {
		return false, errors.New("this BIN is not valid to this brand")
	}

	// check Luhn (mod 10) algorithm
	noSpacesJustDigits :=
		strings.Split(
			strings.Join(
				strings.Split(c.number, " "), ""), "")
	theDigits := []int{}
	var sum int

	for _, d := range noSpacesJustDigits {
		i, err := strconv.Atoi(d)
		if err != nil {
			return false, errors.New(fmt.Sprint("invalid digit: ", d))
		}
		theDigits = append(theDigits, i)
	}

	for i := 0; i < len(theDigits); i++ {
		if i%2 == 0 {
			theDigits[i] *= 2
			if theDigits[i] >= 10 {
				theDigits[i] -= 9
			}
		}
		sum += theDigits[i]
	}

	return sum%10 == 0, nil
}
