package main

import "github.com/orvice/kit/log"

var (
	h      *Hub
	logger log.Logger
	mtr    *monitor
)

func main() {
	InitEnv()
	h = NewHub()
	logger = log.NewDefaultLogger()
	go h.Run()
	go web()
	go netInfoTask()
	mtr = newMonitor()
	go mtr.Daemon()
	go pprof()
	go handleGrpc()
	select {}
}
