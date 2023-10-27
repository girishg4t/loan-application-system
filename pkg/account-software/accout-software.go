package accountsoftware

import "github.com/loan-application-system/pkg/model"

func GetBalanceSheet(username string) []model.Account {
	return []model.Account{
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: 250000,
			AssetsValue:  1234,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: 1150,
			AssetsValue:  5789,
		},
		{
			Year:         2020,
			Month:        10,
			ProfitOrLoss: 2500,
			AssetsValue:  22345,
		},
		{
			Year:         2020,
			Month:        9,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
	}
}
