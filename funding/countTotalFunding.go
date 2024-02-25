package funding

import (
	"fmt"
	"strconv"
	"funding2.0/types"
)

func CountTotalFundingBybit(list []types.BybitSecondItem, symbol string, timeMap *map[string]int64) *types.TotalFundingInDays {
	obj := types.TotalFundingInDays{
		Symbol: symbol,
	}
	var count float64

	for _, v := range list {
		timestamp, err := strconv.ParseInt(v.FundingRateTimestamp, 10, 64)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в int64 CountTotalFundingBybit:", err)
			continue
		}

		fundingNum, err := ConvertToFloat(v.FundingRate)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число CountTotalFundingBybit:", err)
			continue
		}
		
		count += fundingNum

		switch timestamp {
		case (*timeMap)["1"]: 
			obj.OneDay, _ = calculateAndAssignPercent(count, "Bybit")	
		case (*timeMap)["3"]: 
			obj.ThreeDays, _ = calculateAndAssignPercent(count, "Bybit")
		case (*timeMap)["7"]: 
			obj.SevenDays, _ = calculateAndAssignPercent(count, "Bybit")
		case (*timeMap)["14"]: 
			obj.FourteenDays, _ = calculateAndAssignPercent(count, "Bybit")
		case (*timeMap)["30"]: 
			obj.ThirtyDays, _ = calculateAndAssignPercent(count, "Bybit")
			return &obj
		}
	}

	return &obj
}

func CountTotalFundingMexc(list []types.MexcSecondItem, symbol string, timeMap *map[string]int64) *types.TotalFundingInDays {
	obj := types.TotalFundingInDays{
		Symbol: symbol,
	}
	var count float64

	for _, v := range list {
		count += v.FundingRate

		switch v.SettleTime {
		case (*timeMap)["1"]: 
			obj.OneDay, _ = calculateAndAssignPercent(count, "Mexc")
		case (*timeMap)["3"]: 
			obj.ThreeDays, _ = calculateAndAssignPercent(count, "Mexc")
		case (*timeMap)["7"]: 
			obj.SevenDays, _ = calculateAndAssignPercent(count, "Mexc")
		case (*timeMap)["14"]: 
			obj.FourteenDays, _ = calculateAndAssignPercent(count, "Mexc")
		case (*timeMap)["30"]: 
			obj.ThirtyDays, _ = calculateAndAssignPercent(count, "Mexc")
			return &obj
		}
	}

	return &obj
}

func CountTotalFundingKucoin(list []types.KucoinSecondItem, symbol string, timeMap *map[string]int64) *types.TotalFundingInDays {
	obj := types.TotalFundingInDays{
		Symbol: symbol,
	}
	var count float64

	for _, v := range list {
		count += v.FundingRate

		switch v.Timepoint {
		case (*timeMap)["1"]: 
			obj.OneDay, _ = calculateAndAssignPercent(count, "Kucoin")
		case (*timeMap)["3"]: 
			obj.ThreeDays, _ = calculateAndAssignPercent(count, "Kucoin")
		case (*timeMap)["7"]: 
			obj.SevenDays, _ = calculateAndAssignPercent(count, "Kucoin")
		case (*timeMap)["14"]: 
			obj.FourteenDays, _ = calculateAndAssignPercent(count, "Kucoin")
		case (*timeMap)["30"]: 
			obj.ThirtyDays, _ = calculateAndAssignPercent(count, "Kucoin")
			return &obj
		}
	}

	return &obj
}


func CountTotalFundingOkx(list []types.OkxSecondItem, symbol string, timeMap *map[string]int64) *types.TotalFundingInDays {
	obj := types.TotalFundingInDays{
		Symbol: symbol,
	}
	var count float64

	for _, v := range list {
		timestamp, err := strconv.ParseInt(v.FundingTime, 10, 64)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в int64 CountTotalFundingOkx:", err)
			continue
		}

		fundingNum, err := ConvertToFloat(v.FundingRate)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число CountTotalFundingOkx:", err)
			continue
		}
		
		count += fundingNum

		switch timestamp {
		case (*timeMap)["1"]: 
			obj.OneDay, _ = calculateAndAssignPercent(count, "Okx")
		case (*timeMap)["3"]: 
			obj.ThreeDays, _ = calculateAndAssignPercent(count, "Okx")
		case (*timeMap)["7"]: 
			obj.SevenDays, _ = calculateAndAssignPercent(count, "Okx")
		case (*timeMap)["14"]: 
			obj.FourteenDays, _ = calculateAndAssignPercent(count, "Okx")
		case (*timeMap)["30"]: 
			obj.ThirtyDays, _ = calculateAndAssignPercent(count, "Okx")
			return &obj
		}
	}

	return &obj
}

func CountTotalFundingBingx(list []types.BingxSecondItem, symbol string, timeMap *map[string]int64) *types.TotalFundingInDays {
	obj := types.TotalFundingInDays{
		Symbol: symbol,
	}
	var count float64

	for _, v := range list {
		fundingNum, err := ConvertToFloat(v.FundingRate)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число CountTotalFundingBingx:", err)
			continue
		}
		
		count += fundingNum

		switch v.FundingTime {
		case (*timeMap)["1"]: 
			obj.OneDay, _ = calculateAndAssignPercent(count, "BingX")
		case (*timeMap)["3"]: 
			obj.ThreeDays, _ = calculateAndAssignPercent(count, "BingX")
		case (*timeMap)["7"]: 
			obj.SevenDays, _  = calculateAndAssignPercent(count, "BingX")
		case (*timeMap)["14"]: 
			obj.FourteenDays, _ = calculateAndAssignPercent(count, "BingX")
		case (*timeMap)["30"]: 
			obj.ThirtyDays, _ = calculateAndAssignPercent(count, "BingX")
			return &obj
		}
	}

	return &obj
}

func calculateAndAssignPercent(count float64, exchange string) (float64, error) {
    percent, err := ConvertToFloat(fmt.Sprintf("%.4f", count * 100))
    if err != nil {
        fmt.Println("Error converting to float:", exchange, err)
    }
    return percent, err
}
