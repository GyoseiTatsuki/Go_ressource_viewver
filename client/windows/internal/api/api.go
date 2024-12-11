package api

import (
	"Go_ressource_viewver/client/windows/internal/collecte" // Chemin d'importation correct

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/collecte", func(c *gin.Context) {
		stats, err := collecte.CollectStats()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, stats)
	})
}
