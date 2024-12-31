package main

import (
	"net/http" // Pour la gestion des requêtes HTTP

	"github.com/gin-gonic/gin" // Framework Gin pour faciliter la création de l'API HTTP
	"github.com/shirou/gopsutil/cpu"  // Pour récupérer l'utilisation du CPU
	"github.com/shirou/gopsutil/disk" // Pour récupérer l'utilisation du disque
	"github.com/shirou/gopsutil/mem"  // Pour récupérer l'utilisation de la mémoire
)

// Structure représentant les statistiques système
type SystemStats struct {
	CPUUsage    float64 `json:"CPUUsage"`    // Utilisation du CPU en pourcentage
	MemoryUsage float64 `json:"MemoryUsage"` // Utilisation de la mémoire en pourcentage
	DiskUsage   float64 `json:"DiskUsage"`   // Utilisation du disque en pourcentage
}

// Fonction pour récupérer les statistiques système (CPU, mémoire, disque)
func getSystemStats() (SystemStats, error) {
	// Récupérer l'utilisation du CPU en pourcentage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		// Si une erreur survient lors de la récupération des données CPU, on retourne l'erreur
		return SystemStats{}, err
	}

	// Récupérer les statistiques de la mémoire virtuelle
	memStats, err := mem.VirtualMemory()
	if err != nil {
		// Si une erreur survient lors de la récupération des données mémoire, on retourne l'erreur
		return SystemStats{}, err
	}

	// Récupérer les statistiques d'utilisation du disque sur la partition "C:"
	diskStats, err := disk.Usage("C:\\")
	if err != nil {
		// Si une erreur survient lors de la récupération des données du disque, on retourne l'erreur
		return SystemStats{}, err
	}

	// Créer la structure SystemStats avec les données obtenues
	stats := SystemStats{
		CPUUsage:    cpuPercent[0],      // Utilisation du CPU (la première valeur du tableau)
		MemoryUsage: memStats.UsedPercent, // Utilisation de la mémoire (en pourcentage)
		DiskUsage:   diskStats.UsedPercent, // Utilisation du disque (en pourcentage)
	}

	// Retourner les statistiques
	return stats, nil
}

// Handler pour la route /stats de l'API
func statsHandler(c *gin.Context) {
	// Obtenir les statistiques système via la fonction getSystemStats
	stats, err := getSystemStats()
	if err != nil {
		// Si une erreur se produit lors de la récupération des statistiques, on renvoie une erreur HTTP 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Si tout va bien, on renvoie les statistiques sous forme de JSON avec un statut HTTP 200
	c.JSON(http.StatusOK, stats)
}

// Fonction principale qui configure et lance le serveur HTTP
func main() {
	// Créer une instance du routeur Gin par défaut
	r := gin.Default()

	// Définir la route "/stats" qui appelle statsHandler pour renvoyer les statistiques système
	r.GET("/stats", statsHandler)

	// Démarrer le serveur HTTP sur l'adresse 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}
