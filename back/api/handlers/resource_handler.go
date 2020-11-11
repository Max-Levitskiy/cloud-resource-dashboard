package handlers

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers/common/response"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	case http.MethodOptions:
		response.Ok(w)
		break
	default:
		response.NotFound(w)
	}
}

func ResourceCountHandler(c *gin.Context) {
	c.JSON(200, elasticsearch.Client.CountResources())
}

func ResourceDistinctServiceHandler(c *gin.Context) {
	field := strings.Title(c.Param("field"))
	c.JSON(200, elasticsearch.Client.DistinctResourceField(field))
}
