package model

type Account struct {
	Year         int `json:"year"`
	Month        int `json:"month"`
	ProfitOrLoss int `json:"profitOrLoss"`
	AssetsValue  int `json:"assetsValue"`
}

type AssetValue struct {
	Count         int
	SumAssetValue int
	ProfitOrLoss  int
}

type FinalOutput struct {
	Name              string
	EstablishedYear   int
	PreAssessment     int
	SummaryProfitLoss map[int]*AssetValue
}

type UserApplication struct {
	BusinessName    string
	EstablishedYear int
	LoanAmount      int
	Address         string
	AccountProvider string
}
