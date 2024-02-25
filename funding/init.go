package funding

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"funding2.0/types"
)

func Init() ([][]string, map[string]int64, error) {
	bybitUrl := []string{"https://api.bybit.com/v5/market/tickers?category=linear", "https://api.bybit.com/v5/market/funding/history?category=linear&symbol="}
	mexcUrl := []string{"https://contract.mexc.com/api/v1/contract/ticker", "https://contract.mexc.com/api/v1/contract/funding_rate/history?symbol="}
	kucoin := []string{"https://api-futures.kucoin.com/api/v1/contracts/active", "https://api-futures.kucoin.com/api/v1/contract/funding-rates?symbol="}
	okx := []string{"https://www.okx.com/api/v5/public/open-interest?instType=SWAP", "https://www.okx.com/api/v5/public/funding-rate-history?instId="}
	bingx := []string{"https://open-api.bingx.com/openApi/swap/v2/quote/premiumIndex", "https://open-api.bingx.com/openApi/swap/v2/quote/fundingRate?limit=300&symbol="}
	urlsArrays := [][]string{bybitUrl, mexcUrl, kucoin, okx, bingx}

	timestamp, err := getTimestamps()
	if err != nil {
		fmt.Println(err)
		return urlsArrays, nil, err
	}

	return urlsArrays, timestamp, nil
}

var num = map[string]int{
	"OneDay":       2,
	"ThreeDays":    8,
	"SevenDays":    20,
	"FourteenDays": 41,
	"ThirtyDays":   89,
}

func getTimestamps() (map[string]int64, error) {
	// Делаем запрос
	content := DoReq("https://api.bybit.com/v5/market/funding/history?category=linear&symbol=BTCUSDT")
	var response types.BybitResponse
	err := json.Unmarshal([]byte(*content), &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	maps := make(map[string]int64, 5)
	for k, v := range num {
		val, err := strconv.ParseInt(response.Result.List[v].FundingRateTimestamp, 10, 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		maps[k] = val
	}

	timestamp := map[string]int64{"1": maps["OneDay"], "3": maps["ThreeDays"], "7": maps["SevenDays"], "14": maps["FourteenDays"], "30": maps["ThirtyDays"]}
	return timestamp, nil
}

func GetTimestampsKucoin() (map[string]int64, error) {
	// получаем тамштамп 30 дней назад
	nowKucoin := time.Now()
	thenKucoin := nowKucoin.AddDate(0, 0, -31) // вычитаем 30 дней из текущего времени
	thenKucoin = time.Date(thenKucoin.Year(), thenKucoin.Month(), thenKucoin.Day(), 07, 0, 0, 0, thenKucoin.Location())
	timestampKucoin := thenKucoin.UTC().UnixMilli()

	// Делаем запрос
	url := "https://api-futures.kucoin.com/api/v1/contract/funding-rates?symbol=ETHUSDTM&from=" + fmt.Sprint(timestampKucoin) + "&to=" + fmt.Sprint(time.Now().UnixMilli())
	content := DoReq(url)
	var response types.KucoinResponse
	err := json.Unmarshal([]byte(*content), &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	maps := make(map[string]int64, 5)
	for k, v := range num {
		maps[k] = response.Data[v].Timepoint
	}
	timestamp := map[string]int64{"1": maps["OneDay"], "3": maps["ThreeDays"], "7": maps["SevenDays"], "14": maps["FourteenDays"], "30": maps["ThirtyDays"]}
	return timestamp, nil
}
