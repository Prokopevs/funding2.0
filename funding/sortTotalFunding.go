package funding

import (
	// "reflect"
	"reflect"
	"sort"

	"funding2.0/types"
)

func SortTotalFunding(totalFundingSlice []types.TotalFundingInDays, exchangeName string) *types.ExchangeFunding {
	threeDaysArr := sortArray(totalFundingSlice, "ThreeDays")
	sevenDaysArr := sortArray(totalFundingSlice, "SevenDays")
	fourteenDaysArr := sortArray(totalFundingSlice, "FourteenDays")
	thirtyDaysArr :=  sortArray(totalFundingSlice, "ThirtyDays")

	elemsFun := types.ElemsFunding{
		ThreeDays: threeDaysArr,
		SevenDays: sevenDaysArr, 
		FourteenDays: fourteenDaysArr,
		ThirtyDays:	thirtyDaysArr,
	}

	resultData := types.ExchangeFunding{
		Exchange: exchangeName,
		Elems: elemsFun,
	}

	return &resultData
}

func sortArray(arr []types.TotalFundingInDays, field string) []types.FundingElem{

	switch field {
	case "ThreeDays":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].ThreeDays > arr[j].ThreeDays
		})
	case "SevenDays":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].SevenDays > arr[j].SevenDays
		})
	case "FourteenDays":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].FourteenDays > arr[j].FourteenDays
		})
	case "ThirtyDays":
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].ThirtyDays > arr[j].ThirtyDays
		})
	}

	array := make([]types.FundingElem, 0, 7)
	for i, v := range arr {
		if i == 7 {
			break
		}
		value := reflect.ValueOf(v) 
		persent := value.FieldByName(field) 

		elem := types.FundingElem{
			Symbol: v.Symbol,
			FundingPersent: persent.Interface().(float64),
		}

		array = append(array, elem)
	}
	return array
}

