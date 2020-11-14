package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

type ElbScanner struct{}

func (e ElbScanner) Scan(projectId *string, region *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error) {
	list, err := listElb(region, profileName)
	if err == nil {
		if list.LoadBalancers != nil && len(list.LoadBalancers) > 0 {
			tags := getElbTags(region, list.LoadBalancers)
			for _, elb := range list.LoadBalancers {
				r := &model.Resource{
					CloudProvider: clouds.AWS,
					Service:       "elb",
					ProjectId:     projectId,
					Region:        region,
					ResourceId:    elb.DNSName,
					CreationDate:  elb.CreatedTime,
				}
				if elbTags, ok := (*tags)[*elb.LoadBalancerArn]; ok {
					r.Tags = elbTags
				}
				saveCh <- r
			}
		}
	} else {
		errCh <- err
	}
}

func listElb(region *string, profileName *string) (*elbv2.DescribeLoadBalancersOutput, error) {
	svc := elbv2.New(session.Get(region, profileName))
	return svc.DescribeLoadBalancers(&elbv2.DescribeLoadBalancersInput{})
}

func getElbTags(region *string, elbs []*elbv2.LoadBalancer) *map[string]map[string]string {
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

func describeElbTags(region *string, arns []*string, m *map[string]map[string]string) {
	svc := elbv2.New(session.Get(region))
	input := &elbv2.DescribeTagsInput{
		ResourceArns: arns,
	}
	if describeTagsOutput, err := svc.DescribeTags(input); err == nil {
		for _, tagDescription := range describeTagsOutput.TagDescriptions {
			if tags := elbToResourceTags(tagDescription.Tags); tags != nil {
				(*m)[*tagDescription.ResourceArn] = tags
			}

		}
	} else {
		logger.Warn.Println(err)
	}
}

func elbToResourceTags(tags []*elbv2.Tag) map[string]string {
	result := make(map[string]string)
	for _, tag := range tags {
		result[*tag.Key] = *tag.Value
	}
	return result
}
