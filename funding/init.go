package funding

import "time"

func Init() ([][]string, map[string]int64) {
	bybitUrl := []string{"https://api.bybit.com/v5/market/tickers?category=linear", "https://api.bybit.com/v5/market/funding/history?category=linear&symbol="}
	mexcUrl := []string{"https://contract.mexc.com/api/v1/contract/ticker", "https://contract.mexc.com/api/v1/contract/funding_rate/history?symbol="}
	urlsArrays := [][]string{bybitUrl, mexcUrl}

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
