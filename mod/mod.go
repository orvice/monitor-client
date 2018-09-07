package mod

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
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
