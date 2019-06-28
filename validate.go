package cards

import (
	"strconv"
)

// ValidateNumber returns true if `number` satisfy the Luhn algorithm
func ValidateNumber(number string) bool {
	var (
		sum, _  = strconv.Atoi(number[len(number)-1:])
		nDigits = len(number)
		parity  = nDigits % 2
	)

	for i := 0; i < nDigits-1; i++ {
		digit, _ := strconv.Atoi(string(number[i]))
		if (i % 2) == parity {
			digit = digit * 2
		}

		if digit > 9 {
			digit = digit - 9
		}

		sum = sum + digit
	}

	return (sum % 10) == 0
}
