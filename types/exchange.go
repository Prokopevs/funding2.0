package types

type BybitItem struct {
	Symbol string `json:"symbol"`
}
type BybitResponse struct {
	Result struct {
		List []BybitItem `json:"list"`
	} `json:"result"`
}
type BybitSecondItem struct {
	Symbol               string `json:"symbol"`
	FundingRate          string `json:"fundingRate"`
	FundingRateTimestamp string `json:"fundingRateTimestamp"`
}
type BybitSecondResponse struct {
	Result struct {
		List []BybitSecondItem `json:"list"`
	} `json:"result"`
}

type MexcItem struct {
	Symbol string `json:"symbol"`
}
type MexcResponse struct {
	Data []MexcItem `json:"data"`
}
type MexcSecondItem struct {
	Symbol      string  `json:"symbol"`
	FundingRate float64 `json:"fundingRate"`
	SettleTime  int64  `json:"settleTime"`
}

type MexcSecondResponse struct {
	Data struct {
		ResultList []MexcSecondItem `json:"resultList"`
	} `json:"data"`
}
