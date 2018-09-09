package main

import (
	"encoding/json"
	"time"

	"fmt"
	"github.com/orvice/monitor-client/mod"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type monitor struct {
	lastNetStat net.IOCountersStat
}

func newMonitor() *monitor {
	return new(monitor)
}

func (m *monitor) getNetStat(ns []net.IOCountersStat) (net.IOCountersStat, error) {
	for _, n := range ns {
		if n.Name == netInterfaceName {
			return n, nil
		}
	}
	return net.IOCountersStat{}, fmt.Errorf("net interface %s not found", netInterfaceName)
}

func (m *monitor) GetNetSpeed(n net.IOCountersStat) mod.NetSpeed {
	ret := mod.NetSpeed{
		BytesRecv:   n.BytesRecv - m.lastNetStat.BytesRecv,
		BytesSent:   n.BytesSent - m.lastNetStat.BytesSent,
		PacketsRecv: n.PacketsRecv - m.lastNetStat.PacketsRecv,
		PacketsSent: n.PacketsSent - m.lastNetStat.PacketsSent,
	}
	m.lastNetStat = n
	return ret
}

func (m *monitor) GetNetInfo() mod.NetInfo {

	return lastNetInfo
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

	process, err := load.Misc()
	if err != nil {
		logger.Errorf("get misc error: %v", err)
	}

	ns, err := net.IOCounters(true)
	if err != nil {
		logger.Errorf("get net io error: %v ", err)
	}
	stat, err := m.getNetStat(ns)
	speed := m.GetNetSpeed(stat)
	if err != nil {
		logger.Errorf("get net io error: %v ", err)
	}

	cpuTimes, err := cpu.Times(false)
	if err != nil {
		logger.Errorf("get cpu times error: %v ", err)
	}

	cpuCount, _ := cpu.Counts(true)

	diskUsage, err := disk.Usage("/")
	if err != nil {
		logger.Errorf("get disk usage error: %v ", err)
	}

	systemInfo := mod.SystemInfo{
		MemoryStatus: v,
		AvgLoad:      l,
		Process:      process,
		NetSpeed:     speed,
		CpuCount:     cpuCount,
		DiskUsage:    diskUsage,
		NetInfo:      m.GetNetInfo(),
		NetStat:      stat,
	}

	if len(cpuTimes) != 0 {
		systemInfo.CpuTimesStat = cpuTimes[0]
	}

	return systemInfo, nil
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
