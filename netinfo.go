package main

import (
	"github.com/orvice/monitor-client/enum"
	"github.com/orvice/monitor-client/mod"
	"github.com/orvice/monitor-client/utils"
	"time"
)

var lastNetInfo mod.NetInfo

func netInfoTask() {
	for {
		netInfoUpdate()
		time.Sleep(time.Second)
	}
}

func netInfoUpdate() {
	var out mod.NetInfo
	out.Status = enum.ServerStatusOK
	if utils.IsGfwed() {
		out.Status = enum.ServerStatusGFWed
	}
	lastNetInfo = out
}
