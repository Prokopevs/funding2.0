package funding

import (
	"errors"
	"fmt"
	"strconv"
)

func ConvertToFloat(s string) (float64, error) {
	if s == "" {
		return 0, errors.New("empty string provided")
	}
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(s)
		return 0, err
	}
	return n, nil
}
