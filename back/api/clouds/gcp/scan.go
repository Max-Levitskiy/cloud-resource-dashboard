package gcp

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/resources"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/types"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
)

var scanners = []types.Scanner{
	resources.ScannerCloudFunctions{},
}

func FullScan() {
	for _, s := range session.GetSessions() {
		for _, scanner := range scanners {
			var (
				outCh = make(chan []*model.Resource)
				errCh = make(chan error)
			)
			go scanner.Scan(s, outCh, errCh)
			go func() {
				defer close(outCh)
				defer close(errCh)
				select {
				case r := <-outCh:
					logger.Info.Println(r)
					if len(r) > 0 {
						elasticsearch.Client.BulkSave(r)
					}
				case err := <-errCh:
					logger.Error.Println(err)
				}
			}()
		}
	}
}
