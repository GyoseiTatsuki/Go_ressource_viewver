package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func statsHandler(c *gin.Context) {
	stats, err := getSystemStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting stats: %v", err)})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func main() {
	r := gin.Default()

	// Définir la route pour les statistiques
	r.GET("/stats", statsHandler)

	// Démarrer le serveur sur 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}
