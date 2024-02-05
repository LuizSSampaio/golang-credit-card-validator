package validator

import (
	"fmt"
	"strconv"
)

func Validate(cardNumber string) (bool, error) {
	sum := 0
	isSecond := false
	for _, digit := range cardNumber {
		digit, err := strconv.Atoi(fmt.Sprintf("%c", digit))
		if err != nil {
			return false, err
		}

		if isSecond {
			digit *= 2
		}

		sum += digit / 10
		sum += digit % 10

		isSecond = !isSecond
	}

	return sum%10 == 0, nil
}
