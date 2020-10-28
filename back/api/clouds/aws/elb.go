package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/sirupsen/logrus"
)

func scanElb(accountId *string, region string) {
	logrus.Infof("Start scan EBS for %s account %s region", *accountId, region)
	list, err := ListElb(region)
	if err == nil {
		if list.LoadBalancers != nil && len(list.LoadBalancers) > 0 {
			resources := elbToResources(list.LoadBalancers, accountId, &region, getElbTags(region, list.LoadBalancers))
			elasticsearch.Client.BulkSave(resources)
		}
	} else {
		logrus.Error(err)
	}
	logrus.Infof("Scan EBS for %s region finished", region)
}

func ListElb(region string) (*elbv2.DescribeLoadBalancersOutput, error) {
	session := getSession(region)

	svc := elbv2.New(session)

	result, err := svc.DescribeLoadBalancers(&elbv2.DescribeLoadBalancersInput{})
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func elbToResources(elbs []*elbv2.LoadBalancer, accountId *string, region *string, tags *map[string]map[string]string) []*model.Resource {
	var resources = make([]*model.Resource, len(elbs))
	for i, elb := range elbs {

		resources[i] = &model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "elb",
			AccountId:     accountId,
			Region:        region,
			ResourceId:    elb.DNSName,
			CreationDate:  elb.CreatedTime,
		}
		if elbTags, ok := (*tags)[*elb.LoadBalancerArn]; ok {
			resources[i].Tags = elbTags
		}
	}
	return resources
}

func getElbTags(region string, elbs []*elbv2.LoadBalancer) *map[string]map[string]string {
	var arns []*string
	tagsResult := make(map[string]map[string]string)
	for _, elb := range elbs {
		arns = append(arns, elb.LoadBalancerArn)
		if len(arns) == 20 {
			describeElbTags(region, arns, &tagsResult)
			arns = nil
		}
	}
	describeElbTags(region, arns, &tagsResult)
	return &tagsResult
}

func describeElbTags(region string, arns []*string, m *map[string]map[string]string) error {
	svc := elbv2.New(getSession(region))
	input := &elbv2.DescribeTagsInput{
		ResourceArns: arns,
	}
	if describeTagsOutput, err := svc.DescribeTags(input); err == nil {
		for _, tagDescription := range describeTagsOutput.TagDescriptions {
			if tags := elbToResourceTags(tagDescription.Tags); tags != nil {
				(*m)[*tagDescription.ResourceArn] = tags
			}

		}
		return nil
	} else {
		logrus.Error(err)
		return err
	}
}

func elbToResourceTags(tags []*elbv2.Tag) map[string]string {
	result := make(map[string]string)
	for _, tag := range tags {
		result[*tag.Key] = *tag.Value
	}
	return result
}
