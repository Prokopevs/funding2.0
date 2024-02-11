package funding

import (
	"fmt"
	"strconv"
	"funding2.0/types"
)

func CountTotalFunding(list []types.BybitSecondItem, symbol string, timeMap *map[string]int64) *types.TotalFundingInDays {
	obj := types.TotalFundingInDays{
		Symbol: symbol,
	}
	var count float64

	for _, v := range list {
		timestamp, err := strconv.ParseInt(v.FundingRateTimestamp, 10, 64)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в int64:", err)
			continue
		}

		fundingNum := ConvertToFloat(v.FundingRate)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число:", err)
			continue
		}
		
		count += fundingNum

		switch timestamp {
		case (*timeMap)["3"]: 
			obj.ThreeDays = ConvertToFloat(fmt.Sprintf("%.4f", count * 100))
		case (*timeMap)["7"]: 
			obj.SevenDays = ConvertToFloat(fmt.Sprintf("%.4f", count * 100))
		case (*timeMap)["14"]: 
			obj.FourteenDays = ConvertToFloat(fmt.Sprintf("%.4f", count * 100))
		case (*timeMap)["30"]: 
			obj.ThirtyDays = ConvertToFloat(fmt.Sprintf("%.4f", count * 100))
			return &obj
		}
	}

	return &obj
}