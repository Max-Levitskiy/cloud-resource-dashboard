package resources

import (
	"context"
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/parser"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"google.golang.org/api/cloudfunctions/v1"
)

type ScannerCloudFunctions struct {
}

func (ScannerCloudFunctions) Scan(s *session.Session, outCh chan<- []*model.Resource, errCh chan<- error) {
	if service, err := cloudfunctions.NewService(context.Background(), s.GetCredentialsOption()); err == nil {
		parent := fmt.Sprintf("projects/%s/locations/-", s.GetProject().ProjectId)
		if response, err := service.Projects.Locations.Functions.List(parent).Do(); err == nil {
			var resources []*model.Resource
			for _, function := range response.Functions {
				gcpResourceName := parser.ParseName(function.Name)
				resources = append(resources, &model.Resource{
					CloudId:       function.Name,
					CloudProvider: clouds.GCP,
					Service:       "CloudFunctions",
					ProjectId:     &gcpResourceName.ProjectId,
					ResourceId:    &gcpResourceName.ResourceName,
					Region:        &gcpResourceName.Location,
					Tags:          function.Labels,
				})

			}
			outCh <- resources
		} else {
			errCh <- err
		}
	} else {
		errCh <- err
	}
}
