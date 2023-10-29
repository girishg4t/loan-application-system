package rules

import (
	"context"

	"github.com/loan-application-system/pkg/account_software"
	"github.com/loan-application-system/pkg/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("webhook", func() {
	var ()
	BeforeEach(func() {

	})
	AfterEach(func() {

	})
	Context("Verify Rules", func() {
		It("Verify last 12 months get calculated correctly", func() {
			dates := getLastTwelveMonthKey()
			Expect(len(dates)).Should(Equal(12))
		})
		It("Verify report is generated correctly for Business ABC and Account provider XERO", func() {
			as := account_software.NewAccountSoftware()
			bs := as.GetBalanceSheet(context.Background(), model.UserApplication{BusinessName: "ABC", AccountProvider: "XERO"})
			re := NewRuleEngine(bs)
			report := transformBalanceSheet(bs)
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
			as := account_software.NewAccountSoftware()
			bs := as.GetBalanceSheet(context.Background(), model.UserApplication{BusinessName: "XYZ", AccountProvider: "MYOB"})
			report := transformBalanceSheet(bs)
			Expect(report).ShouldNot(Equal(nil))
			Expect(report.AvgAssetValue).To(Equal(63205))
			Expect(report.ProfitOrLoss).To(Equal(-433350))
			Expect(report.YearWiseProfitOrLoss[2022]).To(Equal(-246350))
			Expect(report.YearWiseProfitOrLoss[2023]).To(Equal(-187000))

			re := NewRuleEngine(bs)

			fr := re.RequestOutcome(model.UserApplication{LoanAmount: 60000})
			Expect(fr.PreAssessment).To(Equal(20))

		})
	})
})
