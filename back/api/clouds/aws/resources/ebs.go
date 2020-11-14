package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EbsScanner struct{}

func (e EbsScanner) Scan(projectId *string, region *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error) {
	ebsList, err := listEBS(region, profileName)
	if err == nil {
		for _, volume := range ebsList.Volumes {
			saveCh <- &model.Resource{
				CloudProvider: clouds.AWS,
				Service:       "ebs",
				ProjectId:     projectId,
				Region:        region,
				ResourceId:    volume.VolumeId,
				CreationDate:  volume.CreateTime,
				Tags:          awsToResourceTags(volume.Tags),
			}
		}
	} else {
		errCh <- err
	}
}

func listEBS(region *string, profileName *string) (*ec2.DescribeVolumesOutput, error) {
	svc := ec2.New(session.Get(region, profileName))
	return svc.DescribeVolumes(&ec2.DescribeVolumesInput{})
}
