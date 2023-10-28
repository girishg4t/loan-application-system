package accountsoftware

import "github.com/loan-application-system/pkg/model"

type IAccountSoftware interface {
	GetBalanceSheet(accountProvider string) []model.Account
}

type AccountSoftware struct {
}

func NewAccountSoftware() AccountSoftware {
	return AccountSoftware{}
}

func (a AccountSoftware) GetBalanceSheet(accountProvider string) []model.Account {
	switch accountProvider {
	case model.XERO_ACCOUNTING_PROVIDER:
		{
			return []model.Account{
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
		}

	case model.MYOB_ACCOUNTING_PROVIDER:
		{
			return []model.Account{
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
		}
	default:
		return nil
	}
}
