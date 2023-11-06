package web

import (
	"LANscan/scan"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func printRequest(r *http.Request) {
	today := time.Now()
	fmt.Printf("[%s] Request from IP address %s, URL: %s\n", today.Format("15:04:05 02.01.06"), r.RemoteAddr, r.URL)
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	printRequest(r)
	query := r.URL.Query()
	ipAddresses := query["ip"]
	results := make(map[string]scan.PingResult)

	var wg sync.WaitGroup
	for _, ip := range ipAddresses {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			pingResult := scan.PingDevice(ip, 1000)
			results[ip] = pingResult
		}(ip)
	}

	wg.Wait()

	jsonResponse, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func HandlePort(w http.ResponseWriter, r *http.Request) {
	printRequest(r)
	query := r.URL.Query()
	ipAddresses := query["ip"]
	ports := query["port"]

	results := make(map[string]map[int]scan.PortResult)

	var wg sync.WaitGroup
	for _, ip := range ipAddresses {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			portResults := make(map[int]scan.PortResult)

			if scan.IsValidIP(ip) {
				for _, port := range ports {
					portInt, err := strconv.Atoi(port)
					if err == nil {
						addr := ip + ":" + port
						portResult := scan.PingPort(addr)
						portResults[portInt] = portResult
					}
				}
			}

			results[ip] = portResults
		}(ip)
	}

	wg.Wait()

	jsonResponse, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
