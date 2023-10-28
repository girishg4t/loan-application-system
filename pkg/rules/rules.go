package rules

import (
	"time"

	"github.com/loan-application-system/pkg/model"
)

const (
	DEFAULT_PREASSESSMENT = 20
	FAVORED_PREASSESSMENT = 60
	FULL_PREASSESSMENT    = 100
)

type RuleEngine struct {
	lastYear int
}

func NewRuleEngine() RuleEngine {
	// Assuming loan is applied in current year only
	t := time.Now()
	year := t.Year()
	return RuleEngine{
		lastYear: year - 1,
	}
}

func (r RuleEngine) GetSummary(a model.UserApplication, bs []model.Account) model.FinalOutput {
	report := r.MakeReport(bs)
	s := r.getPnLSummary(report)
	// Apply rules on last 12 months
	p := r.ruleEngine(report, a.LoanAmount)
	return model.FinalOutput{
		Name:              a.BusinessName,
		EstablishedYear:   a.EstablishedYear,
		PreAssessment:     p,
		SummaryProfitLoss: s,
	}
}

func (r RuleEngine) ruleEngine(report map[int]*model.Summary, loanAmount int) int {
	if report[r.lastYear].ProfitOrLoss > 0 {
		avgAssetValue := report[r.lastYear].SumAssetValue / report[r.lastYear].Count
		if avgAssetValue > loanAmount {
			return FULL_PREASSESSMENT
		}
		return FAVORED_PREASSESSMENT
	}
	return DEFAULT_PREASSESSMENT
}

func (r RuleEngine) getPnLSummary(rep map[int]*model.Summary) map[int]int {
	var plSum map[int]int = make(map[int]int)
	for key, value := range rep {
		plSum[key] = value.ProfitOrLoss
	}
	return plSum
}

func (r RuleEngine) MakeReport(bs []model.Account) map[int]*model.Summary {
	var yearAvgAssetValue = make(map[int]*model.Summary)
	for _, val := range bs {
		if yearAvgAssetValue[val.Year] == nil {
			yearAvgAssetValue[val.Year] = &model.Summary{}
		}
		yearAvgAssetValue[val.Year].SumAssetValue = yearAvgAssetValue[val.Year].SumAssetValue + val.AssetsValue
		yearAvgAssetValue[val.Year].Count++
		yearAvgAssetValue[val.Year].ProfitOrLoss = yearAvgAssetValue[val.Year].ProfitOrLoss + val.ProfitOrLoss
	}

	return yearAvgAssetValue
}
