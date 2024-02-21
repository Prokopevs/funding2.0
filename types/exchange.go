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
	SettleTime  int64   `json:"settleTime"`
}

type MexcSecondResponse struct {
	Data struct {
		ResultList []MexcSecondItem `json:"resultList"`
	} `json:"data"`
}

type KucoinItem struct {
	Symbol string `json:"symbol"`
}
type KucoinResponse struct {
	Data []KucoinItem `json:"data"`
}

type KucoinSecondItem struct {
	Symbol      string  `json:"symbol"`
	FundingRate float64 `json:"fundingRate"`
	Timepoint   int64   `json:"timepoint"`
}

type KucoinSecondResponse struct {
	Data []KucoinSecondItem `json:"data"`
}

type OkxItem struct {
	InstId string `json:"instId"`
}
type OkxResponse struct {
	Data []OkxItem `json:"data"`
}
type OkxSecondItem struct {
	InstId      string `json:"instId"`
	FundingRate string `json:"fundingRate"`
	FundingTime string `json:"fundingTime"`
}

type OkxSecondResponse struct {
	Data []OkxSecondItem `json:"data"`
}

type BingxItem struct {
	Symbol string `json:"symbol"`
}
type BingxResponse struct {
	Data []BingxItem `json:"data"`
}
type BingxSecondItem struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64 `json:"fundingTime"`
}

type BingxSecondResponse struct {
	Data []BingxSecondItem `json:"data"`
}
