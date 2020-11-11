package main

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"net/http"
)

func main() {
	runApp()
}

func runApp() {
	http.HandleFunc("/ping", handlers.StatusHandler)
	http.HandleFunc("/resource", handlers.ResourceHandler)
	http.HandleFunc("/resource/count", handlers.ResourceCountHandler)
	http.HandleFunc("/resource/distinct/service", handlers.ResourceDistinctServiceHandler)
	http.HandleFunc("/resource/scan/full", handlers.FullScanHandler)

	logger.Info.Print("Starting api server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error.Fatal(err)
	}
}
