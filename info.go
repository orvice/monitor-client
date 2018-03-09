package main

import (
	"encoding/json"
	"time"

	"github.com/orvice/monitor-client/mod"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type monitor struct {
}

func newMonitor() *monitor {
	return new(monitor)
}

func (m *monitor) GetInfo() (mod.SystemInfo, error) {
	var err error
	v, err := mem.VirtualMemory()
	if err != nil {
		logger.Errorf("get virtual memory error: %v", err)
	}

	l, err := load.Avg()
	if err != nil {
		logger.Errorf("get load error: %v", err)
	}

	n, err := net.IOCounters(true)
	if err != nil {
		logger.Errorf("get net io error: %v ", err, n)
	}

	return mod.SystemInfo{
		MemoryStatus: v,
		AvgLoad:      l,
	}, nil
}

func (m *monitor) SendInfo() error {
	info, err := m.GetInfo()
	if err != nil {
		return err
	}
	b, err := json.Marshal(info)
	if err != nil {
		return err
	}
	h.Broadcast(b)
	return nil
}

func (m *monitor) Daemon() {
	for {
		err := m.SendInfo()
		if err != nil {
		}
		time.Sleep(time.Second)
	}
}
