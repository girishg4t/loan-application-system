package rules

import "github.com/loan-application-system/pkg/model"

type RuleEngine struct {
	balanceSheet []model.Account
}

func NewRuleEngine(bs []model.Account) RuleEngine {
	return RuleEngine{
		balanceSheet: bs,
	}
}

func (r RuleEngine) RuleEngine(year int, loanAmount int) int {
	aav := r.GetAvgAssetValue()
	if aav[year].ProfitOrLoss > 0 {
		avgAssetValue := aav[year].SumAssetValue / aav[year].Count
		if avgAssetValue > loanAmount {
			return 100
		}
		return 60
	}
	return 20
}

func (r RuleEngine) GetAvgAssetValue() map[int]*model.AssetValue {
	var yearAvgAssetValue = make(map[int]*model.AssetValue)
	for _, val := range r.balanceSheet {
		if yearAvgAssetValue[val.Year] == nil {
			yearAvgAssetValue[val.Year] = &model.AssetValue{}
		}
		yearAvgAssetValue[val.Year].SumAssetValue = yearAvgAssetValue[val.Year].SumAssetValue + val.AssetsValue
		yearAvgAssetValue[val.Year].Count++
		yearAvgAssetValue[val.Year].ProfitOrLoss = yearAvgAssetValue[val.Year].ProfitOrLoss + val.ProfitOrLoss
	}

	return yearAvgAssetValue
}
