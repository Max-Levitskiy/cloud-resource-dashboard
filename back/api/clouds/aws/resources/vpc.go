package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	session2 "github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/sirupsen/logrus"
)

func ScanVpc(projectId *string, region string) {
	logrus.Infof("Start Scan VPC for %s account %s region", *projectId, region)
	list, err := listVpc(region)
	if err == nil {
		if list.Vpcs != nil && len(list.Vpcs) > 0 {
			resources := vpcToResources(list.Vpcs, projectId, &region)
			elasticsearch.Client.BulkSave(resources)
		}
	} else {
		logrus.Error(err)
	}
	logrus.Infof("Scan VPC for %s region finished", region)
}

func listVpc(region string) (*ec2.DescribeVpcsOutput, error) {
	session := session2.Get(region)

	svc := ec2.New(session)

	result, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func vpcToResources(vpcs []*ec2.Vpc, projectId *string, region *string) []*model.Resource {
	var resources = make([]*model.Resource, len(vpcs))
	for i, vpc := range vpcs {

		resources[i] = &model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "vpc",
			ProjectId:     projectId,
			Region:        region,
			ResourceId:    vpc.VpcId,
			Tags:          vpcToResourceTags(vpc.Tags),
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
