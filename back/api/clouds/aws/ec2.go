package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func scanEc2(accountId *string, region string) {
	log.Printf("Start scan EC2 for %s account %s region", *accountId, region)
	listS3, err := ListEC2(region)
	if err == nil {
		resources := ec2InstancesToResources(listS3.Reservations, accountId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	log.Printf("Scan EC2 for %s region finished", region)
}

func ListEC2(region string) (*ec2.DescribeInstancesOutput, error) {
	input := &ec2.DescribeInstancesInput{}
	session := getSession(region)

	svc := ec2.New(session)

	result, err := svc.DescribeInstances(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func ec2InstancesToResources(reservations []*ec2.Reservation, accountId *string, region *string) []*model.Resource {
	var resources = make([]*model.Resource, len(reservations))
	for i, reservation := range reservations {
		for _, instance := range reservation.Instances {
			resources[i] = &model.Resource{
				CloudProvider: clouds.AWS,
				Service:       "ec2",
				AccountId:     accountId,
				Region:        region,
				ResourceId:    instance.InstanceId,
				CreationDate:  instance.LaunchTime,
				Tags:          awsToResourceTags(instance.Tags),
			}
		}
	}
	return resources
}
