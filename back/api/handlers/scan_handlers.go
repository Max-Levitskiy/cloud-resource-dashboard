package handlers

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers/common/response"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/scan"
	"net/http"
)

func FullScanHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		go scan.StartFullScan()
		response.Ok(w)
		break
	case http.MethodOptions:
		response.Ok(w)
		break
	default:
		response.NotFound(w)
	}
}
