// Package cron 定时任务
package cron

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/axiaoxin-com/goutils"
	"github.com/axiaoxin-com/investool/datacenter"
	"github.com/axiaoxin-com/investool/models"
	"github.com/axiaoxin-com/logging"
)

// SyncFundManagers 同步基金经理
func SyncFundManagers() {
	ctx := context.Background()
	if !goutils.IsTradingDay() {
		logging.Info(ctx, "Today is not trading day, exit...")
		return
	}
	managers, err := datacenter.EastMoney.FundMangers(ctx, "all", "penavgrowth", "desc")
	if err != nil {
		logging.Error(ctx, "SyncFundManagers error:"+err.Error())
	}
	managers.SortByYieldse()
	if len(managers) != 0 {
		models.FundManagers = managers
	}

	// 更新文件
	b, err := json.MarshalIndent(managers, "", "  ")
	if err != nil {
		logging.Errorf(ctx, "SyncFundManagers json marshal error:", err)
		promSyncError.WithLabelValues("SyncFundManagers").Inc()
		return
	}
	if err := ioutil.WriteFile(models.FundManagersFilename, b, 0666); err != nil {
		logging.Errorf(ctx, "SyncFundManagers WriteFile error:", err)
		promSyncError.WithLabelValues("SyncFundManagers").Inc()
		return
	}
}
