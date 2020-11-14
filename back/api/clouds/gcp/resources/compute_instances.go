package resources

import (
	"context"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/parser"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"google.golang.org/api/compute/v1"
	"time"
)

type ScannerComputeInstances struct {
}

func (ScannerComputeInstances) Scan(s *session.Session, saveCh chan<- *model.Resource, errCh chan<- error) {
	if service, err := compute.NewService(context.Background(), s.GetCredentialsOption()); err == nil {
		projectId := s.GetProject().ProjectId

		if zoneList, err := service.Zones.List(projectId).Do(); err == nil {
			for _, zone := range zoneList.Items {
				if response, err := service.Instances.List(projectId, zone.Name).Do(); err == nil {
					for _, instance := range response.Items {

						region := parser.GetCleanRegionName(zone.Region)
						resource := model.Resource{
							CloudId:       response.Id + "/" + instance.Name,
							CloudProvider: clouds.GCP,
							Service:       "ComputeInstance",
							ProjectId:     &s.GetProject().ProjectId,
							ResourceId:    &instance.Name,
							Region:        &region,
							Tags:          instance.Labels,
						}
						if creationTime, err := time.Parse(time.RFC3339, instance.CreationTimestamp); err == nil {
							resource.CreationDate = &creationTime
						} else {
							logger.Warn.Println(err)
						}
						saveCh <- &resource
					}
				} else {
					errCh <- err
				}
			}
		} else {
			errCh <- err
		}
	} else {
		errCh <- err
	}
}
