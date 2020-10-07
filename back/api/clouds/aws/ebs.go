package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func scanEBS(accountId *string, region string) {
	log.Printf("Start scan EBS for %s account %s region", *accountId, region)
	ebsList, err := ListEBS(region)
	if err == nil {
		resources := ebsInstancesToResources(ebsList.Volumes, accountId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	log.Printf("Scan EBS for %s region finished", region)
}

func ListEBS(region string) (*ec2.DescribeVolumesOutput, error) {
	input := &ec2.DescribeVolumesInput{}
	session := getSession(region)

	svc := ec2.New(session)

	result, err := svc.DescribeVolumes(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func ebsInstancesToResources(volumes []*ec2.Volume, accountId *string, region *string) []model.Resource {
	var resources = make([]model.Resource, len(volumes))
	for i, volume := range volumes {

		resources[i] = model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "ebs",
			AccountId:     accountId,
			Region:        region,
			ResourceId:    volume.VolumeId,
			CreationDate:  volume.CreateTime,
			Tags:          awsToResourceTags(volume.Tags),
		}
	}
	return resources
}
