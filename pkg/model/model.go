package model

type Account struct {
	Year         int `json:"year"`
	Month        int `json:"month"`
	ProfitOrLoss int `json:"profitOrLoss"`
	AssetsValue  int `json:"assetsValue"`
}

type Summary struct {
	Count         int
	SumAssetValue int
	ProfitOrLoss  int
}

type FinalOutput struct {
	Name              string
	EstablishedYear   int
	PreAssessment     int
	SummaryProfitLoss map[int]int
}

type UserApplication struct {
	BusinessName    string
	EstablishedYear int
	LoanAmount      int
	Address         string
	AccountProvider string
}

type FinalOutcome struct {
	Decision       bool `json:"decision"`
	ApprovedAmount int  `json:"approved_amount"`
}
