package handlers

import (
	"net/http"
)

func ResourceHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		//if listS3, err := aws.ListS3("eu-central-1"); err == nil {
		//	resources := aws.s3BucketsToResources(listS3.Buckets)
		//	if marshaled, err := json.Marshal(resources); err == nil {
		//
		//		_, _ = w.Write(marshaled)
		//	} else {
		//		errors.HandleError(err, w, req)
		//	}
		//} else {
		//	errors.HandleError(err, w, req)
		//}
		break
	default:
		w.WriteHeader(404)
	}
}