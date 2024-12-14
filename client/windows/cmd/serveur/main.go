package main

import (
	"Go_ressource_viewver/client/windows/internal/api" // Chemin d'importation correct

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
