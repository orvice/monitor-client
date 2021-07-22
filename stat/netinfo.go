package stat

import (
	"github.com/orvice/monitor-client/enum"
	"github.com/orvice/monitor-client/mod"
	"github.com/orvice/monitor-client/utils"
	"github.com/weeon/log"
	"github.com/weeon/utils/task"
	"time"
)

var lastNetInfo mod.NetInfo

func NetInfoTask() {
	task.NewTaskAndRun("netInfoUpdate", time.Minute, func() error {
		netInfoUpdate()
		return nil
	}, task.SetTaskLogger(log.GetDefault()))
}

func netInfoUpdate() {
	var out mod.NetInfo
	out.Status = enum.ServerStatusOK
	if utils.IsGfwed() {
		out.Status = enum.ServerStatusGFWed
	}
	lastNetInfo = out
}
