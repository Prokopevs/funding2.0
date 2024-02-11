package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"funding2.0/funding"
	"funding2.0/types"
)

func main() {
	urlsArrays, timestamp := funding.Init()

	var wg sync.WaitGroup
	for _, u := range urlsArrays {
		wg.Add(1)
		go func(url []string) {
			defer wg.Done()

			content := funding.DoReq(url[0])
			if content != nil {
				_, err := fillMainSlice(content, url[1], &timestamp)
				fmt.Println(err)
				// if error then not push
			}
		}(u)
	}

	wg.Wait()
}

func fillMainSlice(content *[]byte, secondUrl string, timeMap *map[string]int64) (*types.ExchangeFunding, error) {

	if strings.Contains(secondUrl, "bybit") {
		var response types.BybitResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			// errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			// errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		totalFundingSlice := make([]types.TotalFundingInDays, 0, len(response.Result.List)) 
		var wg sync.WaitGroup

		count := 0
		var mu sync.Mutex
		for _, v := range response.Result.List {

			for _, c := range v.Symbol { // проверяем есть ли -
				if c == '-' {
					continue
				}
			}

			count++
			wg.Add(1)
			go func(obj types.BybitItem, mu *sync.Mutex) {
				defer wg.Done()

				url := secondUrl + obj.Symbol

				content := funding.DoReq(url)
				if content == nil {
					fmt.Println("error", url)
					return 
				}
				
				var data types.BybitSecondResponse
				err := json.Unmarshal([]byte(*content), &data)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				
				totalFund := funding.CountTotalFunding(data.Result.List, obj.Symbol, timeMap) //{symbol: btcusdt, 3: сумма всах фандингов за 3 дня, 7: ...}
				
				// проверить totalFund на наличие
				mu.Lock()
				totalFundingSlice = append(totalFundingSlice, *totalFund)
				mu.Unlock()
			}(v, &mu)

			if count == 20 {
				time.Sleep(800 * time.Millisecond)
				count = 0
			}
		}
		wg.Wait()

		result := funding.SortTotalFunding(totalFundingSlice, "Bybit")

		fmt.Printf("%+v\n", result)
		return result, nil
	}

	// if strings.Contains(secondUrl, "mexc") {
	// 	var response types.MexcResponse
	// 	err := json.Unmarshal([]byte(*content), &response)
	// 	if err != nil {
	// 		// errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
	// 		// errW.ErrorHandler(errStr)
	// 		fmt.Println(err)
	// 		return nil, err
	// 	}

	// 	localSlice := make([]types.Coin, 0, len(response.Data))

	// 	for _, b := range response.Data {
	// 		formatSymbol := strings.ReplaceAll(b.Symbol, "_", "")

	// 		newItem := types.Coin{
	// 			Symbol: formatSymbol,
	// 		}
	// 		localSlice = append(localSlice, newItem)
	// 	}

	// 	obj := &types.FundingItem{Exchange: "Mexc", FundingArr: localSlice}

	// 	return obj, nil
	// }

	return nil, errors.New("fail to find exchange")
}


