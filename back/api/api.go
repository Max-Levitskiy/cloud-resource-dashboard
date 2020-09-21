package main

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlers.StatusHandler)
	http.HandleFunc("/resource", handlers.ResourceHandler)
	http.HandleFunc("/resource/scan/full", handlers.FullScanHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
