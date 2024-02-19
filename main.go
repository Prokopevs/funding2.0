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
	readyData := make([]types.ExchangeFunding, 0, 4) 

	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, u := range urlsArrays {
		wg.Add(1)
		go func(url []string) {
			defer wg.Done()

			content := funding.DoReq(url[0])
			if content != nil {
				result, err := fillMainSlice(content, url[1], &timestamp)
				if err != nil {
					return
				}
				mu.Lock()
				readyData = append(readyData, *result)
				mu.Unlock()
			}
		}(u)
	}

	wg.Wait()

	// fmt.Printf("%+v\n", readyData)
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

			skip := false
			for _, c := range v.Symbol { // проверяем есть ли -
				if c == '-' {
					skip = true
				}
			}
			if skip {
				continue
			}

			count++
			wg.Add(1)
			go func(obj types.BybitItem, mu *sync.Mutex) {
				defer wg.Done()

				url := secondUrl + obj.Symbol

				content := funding.DoReq(url)
				if content == nil {
					fmt.Println("error to do req", url)
					return 
				}
				
				var data types.BybitSecondResponse
				err := json.Unmarshal([]byte(*content), &data)
				if err != nil {
					fmt.Println("error unmarshal", err)
					return
				}
				if len(data.Result.List) == 0 {
					fmt.Println("error: empty data most likely to many request bybit", url)
					return
				}
				
				totalFund := funding.CountTotalFundingBybit(data.Result.List, obj.Symbol, timeMap) //{symbol: btcusdt, 3: сумма всах фандингов за 3 дня, 7: ...}
				mu.Lock()
				totalFundingSlice = append(totalFundingSlice, *totalFund)
				mu.Unlock()
			}(v, &mu)

			if count == 20 {
				time.Sleep(900 * time.Millisecond)
				count = 0
			}
		}
		wg.Wait()

		result := funding.SortTotalFunding(totalFundingSlice, "Bybit")
		fmt.Printf("%+v\n", result)
		return result, nil
	}

	if strings.Contains(secondUrl, "mexc") {
		var response types.MexcResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			// errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			// errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		totalFundingSlice := make([]types.TotalFundingInDays, 0, len(response.Data))
		var wg sync.WaitGroup

		count := 0
		var mu sync.Mutex
		for _, v := range response.Data {

			count++
			wg.Add(1)
			go func(obj types.MexcItem, mu *sync.Mutex) {
				defer wg.Done()

				url := secondUrl + obj.Symbol + "&page_num=1&page_size=200"
	
				content := funding.DoReq(url)
				if content == nil {
					fmt.Println("error to do req", url)
					return 
				}
				
				var data types.MexcSecondResponse
				err := json.Unmarshal([]byte(*content), &data)
				if err != nil {
					fmt.Println("error unmarshal", err)
					return
				}
				if len(data.Data.ResultList) == 0 {
					fmt.Println("error: empty data most likely to many request mexc", url)
					return
				}
				
				totalFund := funding.CountTotalFundingMexc(data.Data.ResultList, obj.Symbol, timeMap) //{symbol: btcusdt, 3: сумма всах фандингов за 3 дня, 7: ...}
				
				mu.Lock()
				totalFundingSlice = append(totalFundingSlice, *totalFund)
				mu.Unlock()
			}(v, &mu)

			if count == 10 {
				time.Sleep(1300 * time.Millisecond)
				count = 0
			}
		}
		wg.Wait()

		result := funding.SortTotalFunding(totalFundingSlice, "Mexc")
		fmt.Printf("%+v\n", result)
		
		return result, nil
	}

	if strings.Contains(secondUrl, "kucoin") {
		timestampKucoin := map[string]int64{"3": funding.CountKucoin(3), "7": funding.CountKucoin(7), "14": funding.CountKucoin(14), 
		"30": funding.CountKucoin(30)}

		var response types.KucoinResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			// errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			// errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		totalFundingSlice := make([]types.TotalFundingInDays, 0, len(response.Data)) 
		var wg sync.WaitGroup

		count := 0
		var mu sync.Mutex
		for _, v := range response.Data {
			if v.Symbol == "XBTMH24" {
				continue
			}
			count++
			wg.Add(1)
			go func(obj types.KucoinItem, mu *sync.Mutex) {
				defer wg.Done()

				url := secondUrl + obj.Symbol + "&from=" + fmt.Sprint(timestampKucoin["30"]) + "&to=" + fmt.Sprint(time.Now().UnixMilli())
		
				content := funding.DoReq(url)
				if content == nil {
					fmt.Println("error to do req", url)
					return 
				}
				
				var data types.KucoinSecondResponse
				err := json.Unmarshal([]byte(*content), &data)
				if err != nil {
					fmt.Println("error unmarshal", err)
					return
				}
				if len(data.Data) == 0 {
					fmt.Println("error: empty data most likely to many request kucoin", url)
					return
				}
				
				totalFund := funding.CountTotalFundingKucoin(data.Data, obj.Symbol, &timestampKucoin) //{symbol: btcusdt, 3: сумма всах фандингов за 3 дня, 7: ...}
				mu.Lock()
				totalFundingSlice = append(totalFundingSlice, *totalFund)
				mu.Unlock()
			}(v, &mu)

			if count == 15 {
				time.Sleep(900 * time.Millisecond)
				count = 0
			}
		}
		wg.Wait()
		result := funding.SortTotalFunding(totalFundingSlice, "Kucoin")
		fmt.Printf("%+v\n", result)
		return result, nil
	}

	if strings.Contains(secondUrl, "okx") {
		var response types.OkxResponse
		err := json.Unmarshal([]byte(*content), &response)
		if err != nil {
			// errStr := fmt.Sprintf("Error Unmarshal data %s\n", err.Error())
			// errW.ErrorHandler(errStr)
			fmt.Println(err)
			return nil, err
		}

		totalFundingSlice := make([]types.TotalFundingInDays, 0, len(response.Data)) 
		var wg sync.WaitGroup

		count := 0
		var mu sync.Mutex
		for _, v := range response.Data {

			count++
			wg.Add(1)
			go func(obj types.OkxItem, mu *sync.Mutex) {
				defer wg.Done()

				url := secondUrl + obj.InstId

				content := funding.DoReq(url)
				if content == nil {
					fmt.Println("error to do req", url)
					return 
				}
				
				var data types.OkxSecondResponse
				err := json.Unmarshal([]byte(*content), &data)
				if err != nil {
					fmt.Println("error unmarshal", err)
					return
				}
				if len(data.Data) == 0 {
					fmt.Println("error: empty data most likely to many request bybit", url)
					return
				}
				
				totalFund := funding.CountTotalFundingOkx(data.Data, obj.InstId, timeMap) //{symbol: btcusdt, 3: сумма всах фандингов за 3 дня, 7: ...}
				fmt.Println(totalFund)
				mu.Lock()
				totalFundingSlice = append(totalFundingSlice, *totalFund)
				mu.Unlock()
			}(v, &mu)

			if count == 20 {
				time.Sleep(900 * time.Millisecond)
				count = 0
			}
		}
		wg.Wait()

		result := funding.SortTotalFunding(totalFundingSlice, "Okx")
		fmt.Printf("%+v\n", result)
		return result, nil
	}

	return nil, errors.New("fail to find exchange")
}


