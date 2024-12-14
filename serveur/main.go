package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Structs pour modéliser les données
type Network struct {
	BytesSent uint64 `json:"BytesSent"`
	BytesRecv uint64 `json:"BytesRecv"`
}

type IPData struct {
	IP          string  `json:"ip"`
	CPUUsage    float64 `json:"CPUUsage"`
	MemoryUsage int     `json:"MemoryUsage"`
	DiskUsage   float64 `json:"DiskUsage"`
	Network     Network `json:"Network"`
}

var (
	ipList    []string
	ipData    []IPData
	dataMutex sync.Mutex
)

func main() {
	// Servir le fichier HTML
	http.Handle("/", http.FileServer(http.Dir("./")))

	// API routes
	http.HandleFunc("/add-ip", addIPHandler)
	http.HandleFunc("/data", dataHandler)

	go updateDataPeriodically()

	fmt.Println("Server started on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

// Handler pour ajouter une IP
func addIPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		IP string `json:"ip"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dataMutex.Lock()
	ipList = append(ipList, requestData.IP)
	dataMutex.Unlock()

	w.WriteHeader(http.StatusOK)
}

// Handler pour récupérer les données
func dataHandler(w http.ResponseWriter, r *http.Request) {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ipData)
}

// Fonction pour mettre à jour les données périodiquement
func updateDataPeriodically() {
	for {
		dataMutex.Lock()
		var newData []IPData
		for _, ip := range ipList {
			url := fmt.Sprintf("http://%s:8080/api/collecte", ip)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Failed to fetch data for IP %s: %v\n", ip, err)
				continue
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Failed to read response for IP %s: %v\n", ip, err)
				resp.Body.Close()
				continue
			}
			resp.Body.Close()

			var data IPData
			if err := json.Unmarshal(body, &data); err != nil {
				fmt.Printf("Failed to parse JSON for IP %s: %v\n", ip, err)
				continue
			}
			data.IP = ip
			newData = append(newData, data)
		}
		ipData = newData
		dataMutex.Unlock()
		time.Sleep(5 * time.Second)
	}
}
