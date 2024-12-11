package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Structure pour stocker les données des clients
type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
	Network     struct {
		BytesSent uint64 `json:"bytes_sent"`
		BytesRecv uint64 `json:"bytes_recv"`
	} `json:"network"`
}

func main() {
	r := gin.Default()

	// Route pour récupérer les données du client
	r.GET("/api/collecte", func(c *gin.Context) {
		// Remplacez l'URL par l'adresse de votre client
		resp, err := http.Get("http://localhost:8080/api/collecte")
		if err != nil {
			log.Println("Error fetching data from client:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from client"})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		// Retourner les données du client
		c.Data(http.StatusOK, "application/json", body)
	})

	// Route pour afficher les données sur une page web
	r.GET("/stats", func(c *gin.Context) {
		// Récupérer les données du client
		resp, err := http.Get("http://localhost:8080/api/collecte")
		if err != nil {
			log.Println("Error fetching data from client:", err)
			c.String(http.StatusInternalServerError, "Failed to fetch data from client")
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			c.String(http.StatusInternalServerError, "Failed to read response")
			return
		}

		// Afficher les données sur une page web
		c.HTML(http.StatusOK, "stats.html", gin.H{
			"data": string(body),
		})
	})

	// Charger le template HTML
	r.LoadHTMLFiles("stats.html")

	// Démarrer le serveur sur le port 8080
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
