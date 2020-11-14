package scan

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
)

var providers = map[string]func(saveCh chan<- *model.Resource, errCh chan<- error){
	clouds.AWS: aws.FullScan,
	clouds.GCP: gcp.FullScan,
}
var errCh = initErrChan()

func StartFullScan() {
	elasticsearch.Client.ClearResourceIndex()
	elasticsearch.Client.CreateIndex()
	for _, f := range providers {
		go f(elasticsearch.BulkSaveCh, errCh)
	}
}
func initErrChan() chan error {
	errCh := make(chan error)
	go func() {
		for {
			err := <-errCh
			logger.Error.Println(err)
		}
	}()
	return errCh
}
