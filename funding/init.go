package funding

import "time"

func Init() ([][]string, map[string]int64) {
	bybitUrl := []string{"https://api.bybit.com/v5/market/tickers?category=linear", "https://api.bybit.com/v5/market/funding/history?category=linear&symbol="}
	mexcUrl := []string{"https://contract.mexc.com/api/v1/contract/ticker", "https://contract.mexc.com/api/v1/contract/funding_rate/history?symbol="}
	kucoin := []string{"https://api-futures.kucoin.com/api/v1/contracts/active", "https://api-futures.kucoin.com/api/v1/contract/funding-rates?symbol="}
	okx := []string{"https://www.okx.com/api/v5/public/open-interest?instType=SWAP", "https://www.okx.com/api/v5/public/funding-rate-history?instId="}
	bingx := []string{"https://open-api.bingx.com/openApi/swap/v2/quote/premiumIndex", "https://open-api.bingx.com/openApi/swap/v2/quote/fundingRate?limit=300&symbol="}
	urlsArrays := [][]string{bybitUrl, mexcUrl, kucoin, okx, bingx}

	timestamp := map[string]int64{"3": countTimestamp(3), "7": countTimestamp(7), "14": countTimestamp(14), "30": countTimestamp(30)}

	return urlsArrays, timestamp
}

func countTimestamp(num int) int64 {
	now := time.Now()               // получение текущего времени
	then := now.AddDate(0, 0, -num) // вычитаем 30 дней из текущего времени
	then = time.Date(then.Year(), then.Month(), then.Day(), 11, 0, 0, 0, then.Location())
	timestamp := then.UTC().UnixMilli()
	return timestamp
}

func CountKucoin(num int) int64 {
	nowKucoin := time.Now()
	thenKucoin := nowKucoin.AddDate(0, 0, -num) // вычитаем 30 дней из текущего времени
	thenKucoin = time.Date(thenKucoin.Year(), thenKucoin.Month(), thenKucoin.Day(), 07, 0, 0, 0, thenKucoin.Location())
	timestampKucoin := thenKucoin.UTC().UnixMilli()
	return timestampKucoin
}
