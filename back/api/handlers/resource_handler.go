package handlers

import (
	"encoding/json"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers/errors"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

func ResourceHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s3, err := aws.ListS3("eu-central-1")
		if err == nil {
			resources := s3BucketsToResources(s3.Buckets)
			marshaled, err := json.Marshal(resources)
			if err == nil {
				w.Write(marshaled)
			} else {
				errors.HandleError(err, w, req)
			}
		} else {
			errors.HandleError(err, w, req)
		}
		break
	}
}

func s3BucketsToResources(buckets []*s3.Bucket) []model.Resource {
	var resources = make([]model.Resource, len(buckets))
	for i, bucket := range buckets {
		resources[i] = model.Resource{
			CloudProvider: clouds.AWS,
			Name:          bucket.Name,
			CreationDate:  bucket.CreationDate,
		}
	}
	return resources
}
