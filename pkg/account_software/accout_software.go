package account_software

import (
	"context"
	"fmt"
	"strings"

	"github.com/loan-application-system/pkg/model"
)

type IAccountSoftware interface {
	GetBalanceSheet(ctx context.Context, app model.UserApplication) []model.Account
}

type AccountSoftware struct {
}

func NewAccountSoftware() AccountSoftware {
	return AccountSoftware{}
}

var dummyBalanceSheet = map[string][]model.Account{
	"ABC-XERO": {
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
			Year:         2023,
			Month:        1,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
	},
	"XYZ-MYOB": {
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
	},
}

func (a AccountSoftware) GetBalanceSheet(ctx context.Context, app model.UserApplication) []model.Account {
	return dummyBalanceSheet[strings.ToUpper(fmt.Sprintf("%s-%s", app.BusinessName, app.AccountProvider))]
}
