package main

import (
	"fmt"

	"github.com/orvice/monitor-client/internal/config"
	"github.com/orvice/monitor-client/internal/hub"
	"github.com/orvice/monitor-client/internal/web"
	"github.com/orvice/monitor-client/stat"
	"github.com/weeon/log"
	"go.uber.org/zap"
)

func main() {
	var err error
	config.InitEnv()
	hub.Logger, err = log.NewLogger("/app/log/monitor-client.log", zap.DebugLevel)
	if err != nil {
		fmt.Println("init logger err", err)
	}

	go stat.NetInfoTask()
	mtr := stat.NewMonitor()
	go mtr.Daemon()

	web.Init()
}
