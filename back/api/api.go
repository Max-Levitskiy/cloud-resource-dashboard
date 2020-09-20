package main

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlers.StatusHandler)
	http.HandleFunc("/resource", handlers.ResourceHandler)
	http.HandleFunc("/resource/scan/full", handlers.FullScanHandler)

	http.ListenAndServe(":8080", nil)
}
