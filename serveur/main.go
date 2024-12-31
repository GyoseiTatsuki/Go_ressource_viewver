package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Stats struct {
	CPUUsage    float64 `json:"CPUUsage"`
	MemoryUsage float64 `json:"MemoryUsage"`
	DiskUsage   float64 `json:"DiskUsage"`
}

func fetchStats(apiURL string) (*Stats, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var stats Stats
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &stats, nil
}

func statsHandler(ipList []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html := `<html>
		<head>
			<title>System Stats</title>
		</head>
		<body>
			<h1>System Statistics</h1>
			<ul>`

		for _, ip := range ipList {
			apiURL := fmt.Sprintf("http://%s:8080/stats", ip)
			stats, err := fetchStats(apiURL)
			if err != nil {
				html += fmt.Sprintf("<li><strong>IP %s:</strong> Error fetching stats (%v)</li>", ip, err)
				continue
			}

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

		html += `</ul></body></html>`

		fmt.Fprint(w, html)
	}
}

func main() {
	var ipList []string
	for {
		var ip string
		fmt.Print("Enter a server IP address (or press Enter to finish): ")
		fmt.Scanln(&ip)
		if ip == "" {
			break
		}
		ipList = append(ipList, ip)
	}

	if len(ipList) == 0 {
		fmt.Println("No IP addresses provided. Exiting.")
		return
	}

	http.HandleFunc("/", statsHandler(ipList))
	fmt.Println("Server running on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
