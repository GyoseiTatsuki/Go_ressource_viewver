package main

import (
	"encoding/json" // Pour la manipulation de données JSON
	"fmt"           // Pour formater les messages d'erreur et les réponses
	"net/http"      // Pour les fonctionnalités HTTP

	"github.com/gin-gonic/gin" // Framework Gin pour créer l'API HTTP
	"github.com/shirou/gopsutil/v3/cpu"  // Pour obtenir l'utilisation du CPU
	"github.com/shirou/gopsutil/v3/disk" // Pour obtenir l'utilisation du disque
	"github.com/shirou/gopsutil/v3/mem"  // Pour obtenir l'utilisation de la mémoire
)

// Structure représentant les statistiques système
type SystemStats struct {
	CPUUsage    float64 `json:"CPUUsage"`    // Utilisation du CPU en pourcentage
	MemoryUsage float64 `json:"MemoryUsage"` // Utilisation de la mémoire en pourcentage
	DiskUsage   float64 `json:"DiskUsage"`   // Utilisation du disque en pourcentage
}

// Fonction pour obtenir les statistiques système (CPU, mémoire, disque)
func getSystemStats() (SystemStats, error) {
	// Obtenir l'utilisation du CPU en pourcentage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		// Si une erreur survient lors de la récupération de l'utilisation du CPU, on retourne une erreur
		return SystemStats{}, err
	}

	// Obtenir les statistiques de la mémoire virtuelle
	memStats, err := mem.VirtualMemory()
	if err != nil {
		// Si une erreur survient lors de la récupération de la mémoire, on retourne une erreur
		return SystemStats{}, err
	}

	// Obtenir l'utilisation du disque sur le répertoire racine
	diskStats, err := disk.Usage("/")
	if err != nil {
		// Si une erreur survient lors de la récupération des statistiques du disque, on retourne une erreur
		return SystemStats{}, err
	}

	// Créer une instance de SystemStats avec les valeurs obtenues
	stats := SystemStats{
		CPUUsage:    cpuPercent[0],      // Prendre la première valeur du tableau retourné pour l'usage CPU
		MemoryUsage: memStats.UsedPercent, // Utilisation de la mémoire en pourcentage
		DiskUsage:   diskStats.UsedPercent, // Utilisation du disque en pourcentage
	}

	// Retourner les statistiques obtenues
	return stats, nil
}

// Gestionnaire de la route /stats pour retourner les statistiques système en JSON
func statsHandler(c *gin.Context) {
	// Obtenir les statistiques système
	stats, err := getSystemStats()
	if err != nil {
		// Si une erreur survient, retourner un code HTTP 500 avec un message d'erreur
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting stats: %v", err)})
		return
	}

	// Si tout va bien, retourner les statistiques en format JSON avec le code HTTP 200
	c.JSON(http.StatusOK, stats)
}

// Fonction principale qui initialise et démarre le serveur HTTP
func main() {
	// Créer une instance du routeur Gin par défaut
	r := gin.Default()

	// Définir la route GET "/stats" qui invoque le gestionnaire statsHandler pour renvoyer les statistiques système
	r.GET("/stats", statsHandler)

	// Démarrer le serveur HTTP sur l'adresse 0.0.0.0 et le port 8080
	r.Run("0.0.0.0:8080")
}
