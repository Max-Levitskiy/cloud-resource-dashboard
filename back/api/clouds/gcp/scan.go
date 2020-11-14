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
	resources.ScannerRun{},
	resources.ScannerComputeInstances{},
}

func FullScan(saveCh chan<- *model.Resource, errCh chan<- error) {
	for _, s := range session.GetSessions() {
		for _, scanner := range scanners {
			var (
				outCh = make(chan []*model.Resource)
			)
			go scanner.Scan(s, saveCh, errCh)
			go func() {
				defer close(outCh)
				select {
				case r := <-outCh:
					logger.Info.Println(r)
					if len(r) > 0 {
						elasticsearch.Client.BulkSave(r)
					}
				}
			}()
		}
	}
}
