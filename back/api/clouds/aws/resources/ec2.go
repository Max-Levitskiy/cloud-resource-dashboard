package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Ec2Scanner struct{}

func (e Ec2Scanner) Scan(projectId *string, region *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error) {
	listS3, err := listEC2(region, profileName)
	if err == nil {
		for _, reservation := range listS3.Reservations {
			for _, instance := range reservation.Instances {
				saveCh <- &model.Resource{
					CloudProvider: clouds.AWS,
					Service:       "ec2",
					ProjectId:     projectId,
					Region:        region,
					ResourceId:    instance.InstanceId,
					CreationDate:  instance.LaunchTime,
					Tags:          awsToResourceTags(instance.Tags),
				}
			}
		}
	} else {
		errCh <- err
	}
}

func listEC2(region *string, profileName *string) (*ec2.DescribeInstancesOutput, error) {
	svc := ec2.New(session.Get(region, profileName))
	return svc.DescribeInstances(&ec2.DescribeInstancesInput{})
}

func awsToResourceTags(tags []*ec2.Tag) map[string]string {
	result := map[string]string{}
	for _, tag := range tags {
		result[*tag.Key] = *tag.Value
	}
	return result
}
