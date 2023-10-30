package account_software

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/loan-application-system/pkg/model"
	"github.com/loan-application-system/pkg/redis"
)

type IAccountSoftware interface {
	GetBalanceSheet(ctx context.Context, app model.UserApplication) []model.Account
}

type AccountSoftware struct {
	redis redis.RedisClient
}

const prefixBalanceSheet = "bs-"
const redisDefaultTTL = 24 * 7 * time.Hour

func NewAccountSoftware(ctx context.Context, cfg redis.Config) AccountSoftware {
	rc, err := redis.NewClient(ctx, cfg)
	if err != nil {
		log.Fatalln("error while initializing redis")
	}
	return AccountSoftware{
		redis: rc,
	}
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

// GetBalanceSheet will return balance sheet based on business name and account provider
func (a AccountSoftware) GetBalanceSheet(ctx context.Context, app model.UserApplication) []model.Account {
	key := strings.ToUpper(fmt.Sprintf("%s-%s", app.BusinessName, app.AccountProvider))
	redisKey := fmt.Sprintf("%s%s", prefixBalanceSheet, key)
	var bs []model.Account = []model.Account{}
	res, err := a.redis.Client.Get(ctx, redisKey).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(res), &bs)
		return bs
	}
	value := dummyBalanceSheet[key]

	byVal, err := json.Marshal(value)
	if err != nil {
		log.Println("unable to marshal balance sheet", key)
	}
	err = a.redis.Client.Set(ctx, redisKey, byVal, redisDefaultTTL).Err()
	if err != nil {
		log.Println("error while calling redis set", key)
	}

	return value
}
