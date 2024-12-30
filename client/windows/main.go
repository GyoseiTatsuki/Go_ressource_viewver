package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Structure pour stocker les données système
type SystemStats struct {
	CPUUsage    float64 `json:"CPUUsage"`
	MemoryUsage float64 `json:"MemoryUsage"`
	DiskUsage   float64 `json:"DiskUsage"`
}

// Fonction pour récupérer les statistiques système
func getSystemStats() (SystemStats, error) {
	// Récupérer l'utilisation CPU
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStats{}, err
	}

	// Récupérer l'utilisation mémoire
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err
	}

	// Récupérer l'utilisation disque
	diskStats, err := disk.Usage("C:\\")
	if err != nil {
		return SystemStats{}, err
	}

	// Créer la structure avec les données
	stats := SystemStats{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: memStats.UsedPercent,
		DiskUsage:   diskStats.UsedPercent,
	}

	return stats, nil
}

// Handler pour l'API
func statsHandler(c *gin.Context) {
	stats, err := getSystemStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func main() {
	r := gin.Default()

	// Définir la route pour les statistiques
	r.GET("/stats", statsHandler)

	// Démarrer le serveur sur le port 8080
	r.Run("0.0.0.0:8080")
}
