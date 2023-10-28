package rules

import (
	"github.com/loan-application-system/pkg/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("webhook", func() {
	var (
		bsA []model.Account
		bsB []model.Account
	)
	BeforeEach(func() {
		bsA = []model.Account{
			{
				Year:         2022,
				Month:        12,
				ProfitOrLoss: 250000,
				AssetsValue:  1234,
			},
			{
				Year:         2022,
				Month:        11,
				ProfitOrLoss: 1150,
				AssetsValue:  5789,
			},
			{
				Year:         2022,
				Month:        10,
				ProfitOrLoss: 2500,
				AssetsValue:  22345,
			},
			{
				Year:         2022,
				Month:        9,
				ProfitOrLoss: -187000,
				AssetsValue:  223452,
			},
		}
		bsB = []model.Account{
			{
				Year:         2022,
				Month:        12,
				ProfitOrLoss: -250000,
				AssetsValue:  1234,
			},
			{
				Year:         2022,
				Month:        11,
				ProfitOrLoss: 1150,
				AssetsValue:  5789,
			},
			{
				Year:         2022,
				Month:        10,
				ProfitOrLoss: 2500,
				AssetsValue:  22345,
			},
			{
				Year:         2022,
				Month:        9,
				ProfitOrLoss: -187000,
				AssetsValue:  223452,
			},
		}
	})
	AfterEach(func() {

	})
	Context("Verify Rules", func() {
		It("Verify report is generated correctly for year 2022 and rules are applied for Business A", func() {
			re := NewRuleEngine()
			report := re.MakeReport(bsA)
			summ := report[2022]
			Expect(summ).ShouldNot(Equal(nil))
			Expect(summ.Count).To(Equal(4))
			Expect(summ.SumAssetValue).To(Equal(252820))
			Expect(summ.ProfitOrLoss).To(Equal(66650))

			preAssessment := re.ruleEngine(report, 260000)
			Expect(preAssessment).To(Equal(60))

			preAssessment = re.ruleEngine(report, 60000)
			Expect(preAssessment).To(Equal(100))

			summ = report[2021]
			Expect(summ).ShouldNot(Equal(nil))
		})
		It("Verify report is generated correctly for year 2022 and rules are applied for Business B", func() {
			re := NewRuleEngine()
			report := re.MakeReport(bsB)
			summ := report[2022]
			Expect(summ).ShouldNot(Equal(nil))
			Expect(summ.Count).To(Equal(4))
			Expect(summ.SumAssetValue).To(Equal(252820))
			Expect(summ.ProfitOrLoss).To(Equal(-433350))

			preAssessment := re.ruleEngine(report, 60000)
			Expect(preAssessment).To(Equal(20))
		})
		It("Verify if Summary is getting generated correctly for business A", func() {
			re := NewRuleEngine()
			report := re.GetSummary(model.UserApplication{
				LoanAmount: 6000,
			}, bsA)
			Expect(report).ShouldNot(Equal(nil))
			Expect(report.SummaryProfitLoss).To(Equal(map[int]int{
				2022: 66650,
			}))
			Expect(report.PreAssessment).To(Equal(100))
		})
		It("Verify if Summary is getting generated correctly for business B", func() {
			re := NewRuleEngine()
			report := re.GetSummary(model.UserApplication{
				LoanAmount: 6000,
			}, bsB)
			Expect(report).ShouldNot(Equal(nil))
			Expect(report.SummaryProfitLoss).To(Equal(map[int]int{
				2022: -433350,
			}))
			Expect(report.PreAssessment).To(Equal(20))
		})
	})
})
