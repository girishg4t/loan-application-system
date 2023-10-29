package rules

import (
	"fmt"
	"time"

	"github.com/loan-application-system/pkg/model"
)

const (
	DEFAULT_PRE_ASSESSMENT = 20
	FAVORED_PRE_ASSESSMENT = 60
	FULL_PRE_ASSESSMENT    = 100
)

type IRuleEngine interface {
	RequestOutcome(a model.UserApplication) model.FinalOutput
}

type RuleEngine struct {
	report       model.Report
	balanceSheet []model.Account
	lastMonths   map[string]*model.TransformedBalanceSheet
}

func NewRuleEngine(bs []model.Account) IRuleEngine {
	lm := getLastTwelveMonthKey()
	r := transformBalanceSheet(bs)
	return RuleEngine{
		lastMonths:   lm,
		balanceSheet: bs,
		report:       r,
	}
}

func (r RuleEngine) RequestOutcome(a model.UserApplication) model.FinalOutput {
	// Apply rules on last 12 months
	p := r.applyRule(a.LoanAmount)

	return model.FinalOutput{
		Name:              a.BusinessName,
		EstablishedYear:   a.EstablishedYear,
		PreAssessment:     p,
		SummaryProfitLoss: r.report.YearWiseProfitOrLoss,
	}
}

func (r RuleEngine) applyRule(loanAmount int) int {
	if r.report.ProfitOrLoss > 0 {
		if r.report.AvgAssetValue > loanAmount {
			return FULL_PRE_ASSESSMENT
		}
		return FAVORED_PRE_ASSESSMENT
	}
	return DEFAULT_PRE_ASSESSMENT
}

func transformBalanceSheet(bs []model.Account) model.Report {
	lastMonths := getLastTwelveMonthKey()
	var key string
	var report = model.Report{
		AvgAssetValue:        0,
		ProfitOrLoss:         0,
		YearWiseProfitOrLoss: map[int]int{},
	}
	var avgCount int
	for _, val := range bs {
		key = fmt.Sprintf("%d-%d", val.Year, val.Month)
		if lastMonths[key] != nil {
			lastMonths[key].AssetValue = val.AssetsValue
			lastMonths[key].ProfitOrLoss = val.ProfitOrLoss

			report.ProfitOrLoss = report.ProfitOrLoss + val.ProfitOrLoss
			report.AvgAssetValue = report.AvgAssetValue + val.AssetsValue
			avgCount++
		}
		report.YearWiseProfitOrLoss[val.Year] = report.YearWiseProfitOrLoss[val.Year] + val.ProfitOrLoss
	}
	report.AvgAssetValue = report.AvgAssetValue / avgCount
	return report
}

func getLastTwelveMonthKey() map[string]*model.TransformedBalanceSheet {
	currentTime := time.Now()
	currentTime = time.Date(currentTime.Year(), currentTime.Month(), 10, 23, 0, 0, 0, time.UTC)

	var lastTran = map[string]*model.TransformedBalanceSheet{}
	for i := 1; i <= 12; i++ {
		date := currentTime.AddDate(0, -i, 0)
		lastTran[fmt.Sprintf("%d-%d", int(date.Year()), int(date.Month()))] = &model.TransformedBalanceSheet{}

	}
	return lastTran
}
