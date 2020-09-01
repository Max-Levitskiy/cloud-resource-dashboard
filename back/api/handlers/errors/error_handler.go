package errors

import "net/http"

func HandleError(err error, w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte(err.Error()))
}
