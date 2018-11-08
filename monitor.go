package main

import (
	"encoding/json"
	"time"

	"bytes"
	"fmt"
	"github.com/orvice/monitor-client/mod"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"net/http"
)

type monitor struct {
	lastNetStat  net.IOCountersStat
	lastNetSpeed mod.NetSpeed
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

func (m *monitor) NetSpeedDaemon() {
	for {
		m.setNetSpeed()
		time.Sleep(time.Second)
	}
}

func (m *monitor) GetNetSpeed() mod.NetSpeed {
	return m.lastNetSpeed
}

func (m *monitor) setNetSpeed() error {
	ns, err := net.IOCounters(true)
	if err != nil {
		logger.Errorf("get net io error: %v ", err)
	}
	n, err := m.getNetStat(ns)
	if err != nil {
		return err
	}
	ret := mod.NetSpeed{
		BytesRecv:   n.BytesRecv - m.lastNetStat.BytesRecv,
		BytesSent:   n.BytesSent - m.lastNetStat.BytesSent,
		PacketsRecv: n.PacketsRecv - m.lastNetStat.PacketsRecv,
		PacketsSent: n.PacketsSent - m.lastNetStat.PacketsSent,
	}
	m.lastNetStat = n
	m.lastNetSpeed = ret
	return nil
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

	speed := m.GetNetSpeed()
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

	m.postStat(b)
	h.Broadcast(b)
	return nil
}

func (m *monitor) postStat(b []byte) {
	if len(postUrl) == 0 {
		return
	}
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("post error: ", err)
		return
	}
	defer resp.Body.Close()
	return
}

func (m *monitor) Daemon() {
	go m.NetSpeedDaemon()
	for {
		err := m.SendInfo()
		if err != nil {
		}
		time.Sleep(time.Second)
	}
}
