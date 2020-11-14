package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type VpcScanner struct{}

func (VpcScanner) Scan(projectId *string, region *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error) {
	if list, err := listVpc(region, profileName); err == nil {
		if list.Vpcs != nil && len(list.Vpcs) > 0 {
			for _, vpc := range list.Vpcs {

				saveCh <- &model.Resource{
					CloudProvider: clouds.AWS,
					Service:       "vpc",
					ProjectId:     projectId,
					Region:        region,
					ResourceId:    vpc.VpcId,
					Tags:          vpcToResourceTags(vpc.Tags),
				}
			}
			resources := vpcToResources(list.Vpcs, projectId, region)
			elasticsearch.Client.BulkSave(resources)
		}
	} else {
		errCh <- err
	}

}

func listVpc(region *string, profileName *string) (*ec2.DescribeVpcsOutput, error) {
	svc := ec2.New(session.Get(region, profileName))

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
