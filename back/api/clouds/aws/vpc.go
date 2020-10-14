package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/sirupsen/logrus"
)

func scanVpc(accountId *string, region string) {
	logrus.Infof("Start scan VPC for %s account %s region", *accountId, region)
	list, err := ListVpc(region)
	if err == nil {
		if list.Vpcs != nil && len(list.Vpcs) > 0 {
			resources := vpcToResources(list.Vpcs, accountId, &region)
			elasticsearch.Client.BulkSave(resources)
		}
	} else {
		logrus.Error(err)
	}
	logrus.Infof("Scan VPC for %s region finished", region)
}

func ListVpc(region string) (*ec2.DescribeVpcsOutput, error) {
	session := getSession(region)

	svc := ec2.New(session)

	result, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func vpcToResources(vpcs []*ec2.Vpc, accountId *string, region *string) []model.Resource {
	var resources = make([]model.Resource, len(vpcs))
	for i, vpc := range vpcs {

		resources[i] = model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "vpc",
			AccountId:     accountId,
			Region:        region,
			ResourceId:    vpc.VpcId,
			Tags:  vpcToResourceTags(vpc.Tags),
		}
	}
	return resources
}

func vpcToResourceTags(tags []*ec2.Tag) map[string]string {
	result := make(map[string]string)
	for _, tag := range tags {
		result[*tag.Key] = *tag.Value
	}
	return result
}
