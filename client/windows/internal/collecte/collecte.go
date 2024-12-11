package collecte

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type SystemStats struct {
	CPUUsage    float64
	MemoryUsage float64
	DiskUsage   float64
	Network     NetworkStats
}

type NetworkStats struct {
	BytesSent uint64
	BytesRecv uint64
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

	netIO, err := net.IOCounters(false)
	if err != nil {
		return SystemStats{}, err
	}

	return SystemStats{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: memInfo.UsedPercent,
		DiskUsage:   diskInfo.UsedPercent,
		Network: NetworkStats{
			BytesSent: netIO[0].BytesSent,
			BytesRecv: netIO[0].BytesRecv,
		},
	}, nil
}
