package handlers

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers/common/response"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		response.Ok(w)
	case http.MethodOptions:
		response.Ok(w)
		break
	default:
		response.NotFound(w)
	}
}
