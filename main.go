package main

import (
	"LANscan/web"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping/", web.HandlePing)
	http.HandleFunc("/port/", web.HandlePort)
	err := http.ListenAndServe(":1010", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
