package response

import (
	"log"
	"net/http"
)

func WriteBytes(w http.ResponseWriter, bytes []byte) {
	setDefaultHeaders(w)
	if _, err := w.Write(bytes); err != nil {
		log.Panic(err)
	}
}

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
