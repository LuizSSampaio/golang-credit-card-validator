package validator

import (
	"fmt"
	"strconv"
)

// TODO: Add validation for cardNumber
func Validate(cardNumber string) (bool, error) {
	sum := 0
	isSecond := false
	for _, digit := range cardNumber {
		digit, _ := strconv.Atoi(fmt.Sprintf("%c", digit))

		if isSecond {
			digit *= 2
		}

		sum += digit / 10
		sum += digit % 10

		isSecond = !isSecond
	}

	return sum%10 == 0, nil
}
