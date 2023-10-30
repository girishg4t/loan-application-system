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
		bsA = []model.Account{{
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
				Year:         2023,
				Month:        1,
				ProfitOrLoss: -187000,
				AssetsValue:  223452,
			}}
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
				Year:         2023,
				Month:        1,
				ProfitOrLoss: -187000,
				AssetsValue:  223452,
			},
		}
	})
	AfterEach(func() {

	})
	Context("Verify Rules", func() {
		It("Verify last 12 months get calculated correctly", func() {
			dates := getLastTwelveMonths()
			Expect(len(dates)).Should(Equal(12))
		})
		It("Verify report is generated correctly for Business ABC and Account provider XERO", func() {
			re := NewRuleEngine(bsA)
			report := transformBalanceSheet(bsA)
			Expect(report).ShouldNot(Equal(nil))
			Expect(report.AvgAssetValue).To(Equal(63205))
			Expect(report.ProfitOrLoss).To(Equal(66650))
			Expect(report.YearWiseProfitOrLoss[2022]).To(Equal(253650))
			Expect(report.YearWiseProfitOrLoss[2023]).To(Equal(-187000))

			fr := re.RequestOutcome(model.UserApplication{LoanAmount: 260000})
			Expect(fr.PreAssessment).To(Equal(60))

			fr = re.RequestOutcome(model.UserApplication{LoanAmount: 60000})
			Expect(fr.PreAssessment).To(Equal(100))

		})
		It("Verify report is generated correctly for Business XYZ and Account provider MYOB", func() {
			report := transformBalanceSheet(bsB)
			Expect(report).ShouldNot(Equal(nil))
			Expect(report.AvgAssetValue).To(Equal(63205))
			Expect(report.ProfitOrLoss).To(Equal(-433350))
			Expect(report.YearWiseProfitOrLoss[2022]).To(Equal(-246350))
			Expect(report.YearWiseProfitOrLoss[2023]).To(Equal(-187000))

			re := NewRuleEngine(bsB)

			fr := re.RequestOutcome(model.UserApplication{LoanAmount: 60000})
			Expect(fr.PreAssessment).To(Equal(20))

		})
	})
})
