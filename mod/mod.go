package mod

import (
	"github.com/shirou/gopsutil/cpu"
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
}

type NetSpeed struct {
	BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
}
