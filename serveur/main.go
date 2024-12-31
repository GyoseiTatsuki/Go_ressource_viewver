package main

import (
	"encoding/json" // Pour manipuler les données JSON
	"fmt"           // Pour les fonctions de formatage et d'affichage
	"io/ioutil"     // Pour lire le corps des réponses HTTP
	"net/http"      // Pour gérer les requêtes HTTP
)

// Structure représentant les statistiques du système
type Stats struct {
	CPUUsage    float64 `json:"CPUUsage"`    // Utilisation du CPU en pourcentage
	MemoryUsage float64 `json:"MemoryUsage"` // Utilisation de la mémoire en pourcentage
	DiskUsage   float64 `json:"DiskUsage"`   // Utilisation du disque en pourcentage
}

// Fonction qui récupère les statistiques depuis l'API d'un serveur
func fetchStats(apiURL string) (*Stats, error) {
	// Effectue une requête HTTP GET pour obtenir les statistiques depuis l'URL donnée
	resp, err := http.Get(apiURL)
	if err != nil {
		// Si une erreur se produit lors de la requête, on retourne l'erreur
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}
	defer resp.Body.Close() // S'assurer de fermer le corps de la réponse après utilisation

	// Si le code de statut de la réponse n'est pas 200 (OK), on retourne une erreur
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Lire tout le corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Si la lecture échoue, on retourne une erreur
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Décode le corps JSON dans la structure Stats
	var stats Stats
	if err := json.Unmarshal(body, &stats); err != nil {
		// Si la désérialisation échoue, on retourne une erreur
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Retourne les statistiques obtenues
	return &stats, nil
}

// Fonction qui retourne un handler HTTP pour afficher les statistiques de plusieurs serveurs
func statsHandler(ipList []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Structure de base de la page HTML
		html := `<html>
		<head>
			<title>System Stats</title>
			<script type="text/javascript">
				setTimeout(function(){
					location.reload(); // Rafraîchit la page toutes les 1000 ms (1 seconde)
				}, 1000); 
			</script>
		</head>
		<body>
			<h1>System Statistics</h1>
			<ul>`

		// Parcourt la liste des IPs pour récupérer et afficher les statistiques de chaque serveur
		for _, ip := range ipList {
			// Génère l'URL de l'API pour chaque serveur en fonction de son IP
			apiURL := fmt.Sprintf("http://%s:8080/stats", ip)
			// Récupère les statistiques via la fonction fetchStats
			stats, err := fetchStats(apiURL)
			if err != nil {
				// Si une erreur se produit, on affiche un message d'erreur pour cet IP
				html += fmt.Sprintf("<li><strong>IP %s:</strong> Error fetching stats (%v)</li>", ip, err)
				continue // Passe à l'IP suivante
			}

			// Ajoute les statistiques du serveur dans la page HTML
			html += fmt.Sprintf(`
				<li>
					<strong>IP %s:</strong>
					<ul>
						<li>CPU Usage: %.2f%%</li>
						<li>Memory Usage: %.2f%%</li>
						<li>Disk Usage: %.2f%%</li>
					</ul>
				</li>
			`, ip, stats.CPUUsage, stats.MemoryUsage, stats.DiskUsage)
		}

		// Ferme la liste et la page HTML
		html += `</ul></body></html>`

		// Envoie la page HTML à l'utilisateur
		fmt.Fprint(w, html)
	}
}

// Fonction principale qui initialise le serveur et récupère les IPs à surveiller
func main() {
	var ipList []string
	// Demande à l'utilisateur de saisir des adresses IP des serveurs à surveiller
	for {
		var ip string
		fmt.Print("Enter a server IP address (or press Enter to finish): ")
		fmt.Scanln(&ip)
		if ip == "" {
			// Si l'utilisateur appuie sur Enter sans entrer d'IP, on termine la saisie
			break
		}
		ipList = append(ipList, ip) // Ajoute l'IP à la liste
	}

	// Si aucune IP n'a été fournie, on arrête l'exécution
	if len(ipList) == 0 {
		fmt.Println("No IP addresses provided. Exiting.")
		return
	}

	// Définit le handler HTTP pour la route racine ("/") et démarre le serveur
	http.HandleFunc("/", statsHandler(ipList))
	fmt.Println("Server running on http://0.0.0.0:8081") // Affiche où le serveur écoute
	http.ListenAndServe("0.0.0.0:8081", nil) // Démarre le serveur sur le port 8081
}
