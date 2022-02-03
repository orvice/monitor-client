package mod

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SystemInfo struct {
	MemoryStatus *mem.VirtualMemoryStat `json:"memory_status"`
	AvgLoad      *load.AvgStat          `json:"avg_load"`
	Process      *load.MiscStat         `json:"process"`
	NetSpeed     NetSpeed               `json:"net_speed"`
	CpuCount     int                    `json:"cpu_count"`
	CpuTimesStat cpu.TimesStat          `json:"cpu_times_stat"`
	DiskUsage    *disk.UsageStat        `json:"disk_usage"`
	NetInfo      NetInfo                `json:"net_info"`
	NetStat      net.IOCountersStat     `json:"net_stat"`
}

type NetSpeed struct {
	BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
}

type NetInfo struct {
	Status int32 `json:"status"`
}

type NodeStat struct {
	NodeID   int32      `json:"node_id"`
	NodeCode string     `json:"node_code"`
	NodeName string     `json:"node_name"`
	Stat     SystemInfo `json:"stat"`
	Time     int64      `json:"time"`
}
