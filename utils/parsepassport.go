package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ParsePassport(passport string) (int, int, error) {
	parts := strings.Split(passport, " ")
	if len(parts) != 2 {
		return 0, 0, errors.New("неправильный формат паспорта")
	}

	series, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("неправильная серия паспорта")
	}

	number, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, errors.New("неправильный номер паспорта")
	}
	return series, number, nil
}
