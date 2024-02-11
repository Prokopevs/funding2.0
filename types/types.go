package types

type FundingItem struct {
	Exchange   string
	FundingArr []Coin
}

type SuitableCoin struct {
	Exchange        string  `json:"exchange"`
	Symbol          string  `json:"symbol"`
	BidPrice        string  `json:"bidPrice"`
	AskPrice        string  `json:"askPrice"`
	FundingRate     float64 `json:"fundingRate"`
	NextFundingTime string  `json:"nextFundingTime"`
}

type Coin struct {
	Symbol string `json:"symbol"`
}

type FundingCoin struct {
	Symbol string         `json:"symbol"`
	Elems  []SuitableCoin `json:"elems"`
}

type TotalFundingInDays struct {
	Symbol       string  `json:"symbol"`
	ThreeDays    float64 `json:"threeDays"`
	SevenDays    float64 `json:"sevenDays"`
	FourteenDays float64 `json:"fourteenDays"`
	ThirtyDays   float64 `json:"thirtyDays"`
}

// ----------------------------- //

type FundingElem struct {
	Symbol string `json:"symbol"`
	FundingPersent float64 `json:"fundingPersent"`
}

type ElemsFunding struct {
	ThreeDays    []FundingElem `json:"threeDays"`
	SevenDays    []FundingElem `json:"sevenDays"`
	FourteenDays []FundingElem `json:"fourteenDays"`
	ThirtyDays   []FundingElem `json:"thirtyDays"`
}

type ExchangeFunding struct {
	Exchange string           `json:"exchange"`
	Elems    ElemsFunding `json:"elems"`
}
