package helper

import (
	"errors"
	"strconv"
)

func ConvertInt(num string) (int, error) {
	res, err := strconv.Atoi(num)
	if err != nil {
		return 0, errors.New("value must be numeric")
	}

	return res, nil
}
