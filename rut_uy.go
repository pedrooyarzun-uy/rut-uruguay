package rut_uy

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	res := ValidateRUT("220084260014")
	fmt.Println(res)
}

func ValidateRUT(rut string) bool {

	//Regex for check only numbers
	re := regexp.MustCompile("^[0-9]+$")

	if len(rut) != 12 || !re.MatchString(rut) {
		return false
	}

	lastDigit, err := strconv.Atoi(string(rut[len(rut)-1]))

	if err != nil {
		panic(err)
	}

	total := 0
	factor := 2

	for i := 10; i >= 0; i-- {
		n, err := strconv.Atoi(string(rut[i]))

		if err != nil {
			panic(err)
		}

		total += factor * n

		if factor == 9 {
			factor = 2
		} else {
			factor += 1
		}
	}

	controlDigit := 11 - (total % 11)

	if controlDigit == 11 {
		controlDigit = 0
	} else if controlDigit == 10 {
		controlDigit = 1
	}

	if lastDigit != controlDigit {
		return false
	}

	return true
}
