package stat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/orvice/monitor-client/enum"
	"github.com/orvice/monitor-client/internal/config"
	"github.com/orvice/monitor-client/mod"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/weeon/log"
)

type Monitor struct {
	lastNetStat  net.IOCountersStat
	lastNetSpeed mod.NetSpeed
}

func NewMonitor() *Monitor {
	return new(Monitor)
}

func (m *Monitor) getNetStat(ns []net.IOCountersStat) (net.IOCountersStat, error) {
	for _, n := range ns {
		if n.Name == config.NetInterfaceName {
			return n, nil
		}
	}
	return net.IOCountersStat{}, fmt.Errorf("net interface %s not found", config.NetInterfaceName)
}

func (m *Monitor) NetSpeedDaemon() {
	d := time.Second
	tick := time.NewTicker(d)
	for {
		select {
		case <-tick.C:
			_ = m.setNetSpeed()
			tick.Reset(d)
		}
	}
}

func (m *Monitor) GetNetSpeed() mod.NetSpeed {
	return m.lastNetSpeed
}

func (m *Monitor) setNetSpeed() error {
	ns, err := net.IOCounters(true)
	if err != nil {
		log.Errorf("get net io error: %v ", err)
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

func (m *Monitor) GetNetInfo() mod.NetInfo {

	return lastNetInfo
}

func (m *Monitor) GetInfo() (mod.SystemInfo, error) {
	var err error
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Errorf("get virtual memory error: %v", err)
	}

	l, err := load.Avg()
	if err != nil {
		log.Errorf("get load error: %v", err)
	}

	process, err := load.Misc()
	if err != nil {
		log.Errorf("get misc error: %v", err)
	}

	ns, err := net.IOCounters(true)
	if err != nil {
		log.Errorf("get net io error: %v ", err)
	}
	stat, err := m.getNetStat(ns)

	speed := m.GetNetSpeed()
	if err != nil {
		log.Errorf("get net io error: %v ", err)
	}

	cpuTimes, err := cpu.Times(false)
	if err != nil {
		log.Errorf("get cpu times error: %v ", err)
	}

	cpuCount, _ := cpu.Counts(true)

	diskUsage, err := disk.Usage("/")
	if err != nil {
		log.Errorf("get disk usage error: %v ", err)
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

func (m *Monitor) SendInfo() error {
	info, err := m.GetInfo()
	if err != nil {
		return err
	}
	b, err := json.Marshal(info)
	if err != nil {
		return err
	}

	m.postStat(b)
	return nil
}

func (m *Monitor) postStat(b []byte) {
	if len(config.PostUrl) == 0 {
		return
	}
	req, err := http.NewRequest("POST", config.PostUrl, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(enum.PostKey, config.PostKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("post error: ", err)
		return
	}
	defer resp.Body.Close()
	return
}

func (m *Monitor) Daemon() {
	go m.NetSpeedDaemon()
	for {
		err := m.SendInfo()
		if err != nil {
		}
		time.Sleep(time.Second)
	}
}
