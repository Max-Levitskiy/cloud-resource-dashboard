package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EbsScanner struct {
}

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
	input := &ec2.DescribeVolumesInput{}
	s := session.Get(*region, profileName)

	svc := ec2.New(s)

	result, err := svc.DescribeVolumes(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}
