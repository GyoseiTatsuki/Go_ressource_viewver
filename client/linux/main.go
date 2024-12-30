package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsage    float64 `json:"CPUUsage"`
	MemoryUsage float64 `json:"MemoryUsage"`
	DiskUsage   float64 `json:"DiskUsage"`
}

func getSystemStats() (SystemStats, error) {
	// CPU Usage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStats{}, err
	}

	// Memory Usage
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err
	}

	// Disk Usage
	diskStats, err := disk.Usage("/")
	if err != nil {
		return SystemStats{}, err
	}

	// Populate the struct
	stats := SystemStats{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: memStats.UsedPercent,
		DiskUsage:   diskStats.UsedPercent,
	}

	return stats, nil
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting stats: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	http.HandleFunc("/api/collecte", statsHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
