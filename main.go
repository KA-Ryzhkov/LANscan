package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping/", handlePing)
	err := http.ListenAndServe(":1010", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
