package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	session2 "github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func ScanEc2(projectId *string, region string) {
	log.Printf("Start scan EC2 for %s account %s region", *projectId, region)
	listS3, err := listEC2(region)
	if err == nil {
		resources := ec2InstancesToResources(listS3.Reservations, projectId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	log.Printf("Scan EC2 for %s region finished", region)
}

func listEC2(region string) (*ec2.DescribeInstancesOutput, error) {
	input := &ec2.DescribeInstancesInput{}
	session := session2.Get(region)

	svc := ec2.New(session)

	result, err := svc.DescribeInstances(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func ec2InstancesToResources(reservations []*ec2.Reservation, projectId *string, region *string) []*model.Resource {
	var resources = make([]*model.Resource, len(reservations))
	for i, reservation := range reservations {
		for _, instance := range reservation.Instances {
			resources[i] = &model.Resource{
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
	return resources
}

func awsToResourceTags(tags []*ec2.Tag) map[string]string {
	result := map[string]string{}
	for _, tag := range tags {
		result[*tag.Key] = *tag.Value
	}
	return result
}
