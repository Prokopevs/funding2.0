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
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingRateTimestamp string `json:"fundingRateTimestamp"`
}

type BybitSecondResponse struct {
	Result struct {
		List []BybitSecondItem `json:"list"`
	} `json:"result"`
}

type MexcItem struct {
	Symbol      string  `json:"symbol"`
	Bid1        float64 `json:"bid1"`
	Ask1        float64 `json:"ask1"`
	FundingRate float64 `json:"fundingRate"`
}
type MexcResponse struct {
	Data []MexcItem `json:"data"`
}
