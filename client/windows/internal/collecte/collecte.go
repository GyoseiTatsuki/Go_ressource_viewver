package collecte

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type SystemStats struct {
	CPUUsage    float64
	MemoryUsage float64
	DiskUsage   float64
}

func CollectStats() (SystemStats, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStats{}, err
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err
	}

	diskInfo, err := disk.Usage("/")
	if err != nil {
		return SystemStats{}, err
	}

	return SystemStats{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: memInfo.UsedPercent,
		DiskUsage:   diskInfo.UsedPercent,
	}, nil
}
