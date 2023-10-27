package decisionengine

import (
	"github.com/loan-application-system/pkg/model"
	"github.com/loan-application-system/pkg/rules"
)

func MakeDecision(a model.UserApplication, bs []model.Account, year int) model.FinalOutput {
	r := rules.NewRuleEngine(bs)
	av := r.GetAvgAssetValue()
	pav := r.RuleEngine(year, a.LoanAmount)
	return model.FinalOutput{
		Name:              "",
		EstablishedYear:   0,
		PreAssessment:     pav,
		SummaryProfitLoss: av,
	}
}
