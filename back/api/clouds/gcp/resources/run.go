package resources

import (
	"context"
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"google.golang.org/api/run/v1"
	"time"
)

type ScannerRun struct {
}

func (ScannerRun) Scan(s *session.Session, saveCh chan<- *model.Resource, errCh chan<- error) {
	if service, err := run.NewService(context.Background(), s.GetCredentialsOption()); err == nil {
		parent := fmt.Sprintf("projects/%s/locations/-", s.GetProject().ProjectId)
		if response, err := service.Projects.Locations.Services.List(parent).Do(); err == nil {
			for _, item := range response.Items {
				region := item.Metadata.Labels["cloud.googleapis.com/location"]
				resource := model.Resource{
					CloudId:       item.Metadata.SelfLink,
					CloudProvider: clouds.GCP,
					Service:       "Run",
					ProjectId:     &s.GetProject().ProjectId,
					ResourceId:    &item.Metadata.Name,
					Region:        &region,
					Tags:          item.Metadata.Labels,
				}
				if creationTime, err := time.Parse(time.RFC3339, item.Metadata.CreationTimestamp); err == nil {
					resource.CreationDate = &creationTime
				} else {
					logger.Warn.Println(err)
				}
				saveCh <- &resource
			}
		} else {
			errCh <- err
		}
	} else {
		errCh <- err
	}
}
