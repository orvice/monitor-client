package mod

import (
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	MemoryStatus *mem.VirtualMemoryStat `json:"memory_status"`
	AvgLoad      *load.AvgStat          `json:"avg_load"`
}
