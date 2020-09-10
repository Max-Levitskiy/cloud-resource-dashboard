package handlers

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/scan"
	"net/http"
)

func FullScanHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		go scan.StartFullScan()
		w.WriteHeader(200)
		break
	default:
		w.WriteHeader(404)
	}
}
