package main

import "github.com/orvice/kit/log"

var (
	h      *Hub
	logger log.Logger

	netInterfaceName string
)

func main() {
	h = NewHub()
	logger = log.NewDefaultLogger()
	go h.Run()
	go web()
	monitor := newMonitor()
	go monitor.Daemon()
	select {}
}
