package funding

import (
	"fmt"
	"strconv"
)

func ConvertToFloat(s string) float64 {
	if s == "" {
		return 0
	}
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(s)
		return 0
	}
	return n
}