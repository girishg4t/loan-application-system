package model

type Account struct {
	Year         int `json:"year"`
	Month        int `json:"month"`
	ProfitOrLoss int `json:"profitOrLoss"`
	AssetsValue  int `json:"assetsValue"`
}
type Report struct {
	AvgAssetValue        int
	ProfitOrLoss         int
	YearWiseProfitOrLoss map[int]int
}

type TransformedBalanceSheet struct {
	AssetValue   int
	ProfitOrLoss int
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
