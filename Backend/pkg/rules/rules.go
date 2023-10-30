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
}

func NewRuleEngine(bs []model.Account) IRuleEngine {
	r := transformBalanceSheet(bs)
	return RuleEngine{
		balanceSheet: bs,
		report:       r,
	}
}

// RequestOutcome will take input as user application and return the final output with correct pre assessment value
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

// applyRule calculate pre assessment value according to rules
func (r RuleEngine) applyRule(loanAmount int) int {
	if r.report.ProfitOrLoss > 0 {
		if r.report.AvgAssetValue > loanAmount {
			return FULL_PRE_ASSESSMENT
		}
		return FAVORED_PRE_ASSESSMENT
	}
	return DEFAULT_PRE_ASSESSMENT
}

// transformBalanceSheet convert's the balance sheet to a report for rules calculation
// it take's last 12 months data and calculate Average Asset value, profitorloss for last 12 months and Year wise profitorloss
// here average is calculated based on the data which is present for last 12 months,
// if the data is available for 10 months then the Average is calculated for that many months.
func transformBalanceSheet(bs []model.Account) model.Report {
	last12Months := getLastTwelveMonths()
	var report = model.Report{
		AvgAssetValue:        0,
		ProfitOrLoss:         0,
		YearWiseProfitOrLoss: map[int]int{},
	}
	var avgCount int
	var key string
	for _, val := range bs {
		key = fmt.Sprintf("%d-%d", val.Year, val.Month)
		if last12Months[key] != nil {
			report.ProfitOrLoss = report.ProfitOrLoss + val.ProfitOrLoss
			report.AvgAssetValue = report.AvgAssetValue + val.AssetsValue
			avgCount++
		}
		report.YearWiseProfitOrLoss[val.Year] = report.YearWiseProfitOrLoss[val.Year] + val.ProfitOrLoss
	}
	report.AvgAssetValue = report.AvgAssetValue / avgCount
	return report
}

// getLastTwelveMonths get's the data for last 12 months
// Here date as taken 10 of current month, this is because to avoid months calculation, sine some month as 29 days
// Year-Month wise data get stored for last 12 months only
func getLastTwelveMonths() map[string]*model.TransformedBalanceSheet {
	currentTime := time.Now()
	currentTime = time.Date(currentTime.Year(), currentTime.Month(), 10, 23, 0, 0, 0, time.UTC)

	var lastTran = map[string]*model.TransformedBalanceSheet{}
	for i := 1; i <= 12; i++ {
		date := currentTime.AddDate(0, -i, 0)
		lastTran[fmt.Sprintf("%d-%d", int(date.Year()), int(date.Month()))] = &model.TransformedBalanceSheet{}

	}
	return lastTran
}
