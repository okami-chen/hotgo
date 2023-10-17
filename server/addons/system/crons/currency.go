package crons

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/addons/system/dao"
	"hotgo/internal/library/cron"
	"net/http"
)

// 定时任务.
// 插件中的定时任务可以统一在这里注册和处理

func init() {
	cron.Register(Currency)
}

type ExchangeRateResponse struct {
	Result string             `json:"result"`
	Rates  map[string]float64 `json:"conversion_rates"`
}

var Currency = &CCurrency{name: "currency"}

type CCurrency struct {
	name string
}

func (c *CCurrency) GetName() string {
	return c.name
}

func (c *CCurrency) Execute(ctx context.Context) {
	url := "https://v6.exchangerate-api.com/v6/40f5eda9d1b1283846ee789c/latest/USD"

	response, err := http.Get(url)
	if err != nil {
		cron.Logger().Errorf(ctx, "Error fetching data: %v", err.Error())
		return
	}
	defer response.Body.Close()

	var exchangeRates ExchangeRateResponse
	if err := json.NewDecoder(response.Body).Decode(&exchangeRates); err != nil {
		cron.Logger().Errorf(ctx, "Error decoding response: %v", err.Error())
		return
	}

	for code, rate := range exchangeRates.Rates {
		_, e := dao.Currency.Ctx(ctx).Where("code = ?", code).Update(g.Map{"rate": rate})
		if e != nil {
			cron.Logger().Errorf(ctx, "更新汇率: %v -> %v ")
			continue
		}
		cron.Logger().Infof(ctx, "更新汇率: %v -> %v", code, rate)
	}
}
