package response

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/transform"
	"log"
	"net/http"
)

func WriteAsJson(w http.ResponseWriter, value interface{}) {
	WriteBytes(w, transform.ToJson(value))
}

func WriteBytes(w http.ResponseWriter, bytes []byte) {
	setDefaultHeaders(w)
	if _, err := w.Write(bytes); err != nil {
		log.Panic(err)
	}
}

func Ok(w http.ResponseWriter) {
	WriteBytes(w, []byte(`{"status": "ok"}`))
}

func Status(w http.ResponseWriter, status int) {
	setDefaultHeaders(w)
	w.WriteHeader(status)
}

func NotFound(w http.ResponseWriter) {
	Status(w, http.StatusNotFound)
}

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
