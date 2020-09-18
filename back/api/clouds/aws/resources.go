package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
)

var awsRegions = getAwsRegions()

func FullScan() {
	log.Println("Start AWS scan")
	accountId := getAccountId()
	for _, region := range awsRegions {
		go scanS3(accountId, region)
		go scanEc2(accountId, region)
	}
}

func scanS3(accountId *string, region string) {
	log.Printf("Start scan S3 for %s region", region)
	listS3, err := ListS3(region)
	if err == nil {
		resources := s3BucketsToResources(listS3.Buckets, accountId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	log.Printf("Scan S3 for %s region finished", region)
}

func scanEc2(accountId *string, region string) {
	log.Printf("Start scan EC2 for %p account %s region", accountId, region)
	listS3, err := ListEC2(region)
	if err == nil {
		resources := ec2InstancesToResources(listS3.Reservations, accountId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	log.Printf("Scan EC2 for %s region finished", region)
}

func getAccountId() *string {
	s := sts.New(getSession("us-east-1"))
	identity, err := s.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err == nil {
		return identity.Account
	} else {
		log.Panic(err)
		return nil
	}
}

func ListS3(region string) (*s3.ListBucketsOutput, error) {
	input := &s3.ListBucketsInput{}
	session := getSession(region)

	svc := s3.New(session)

	result, err := svc.ListBuckets(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
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

func s3BucketsToResources(buckets []*s3.Bucket, accountId *string, region *string) []model.Resource {
	var resources = make([]model.Resource, len(buckets))
	for i, bucket := range buckets {
		resources[i] = model.Resource{
			CloudProvider: clouds.AWS,
			ResourceType:  "s3",
			AccountId:     accountId,
			Region:        region,
			Name:          bucket.Name,
			CreationDate:  bucket.CreationDate,
		}
	}
	return resources
}

func ec2InstancesToResources(reservations []*ec2.Reservation, accountId *string, region *string) []model.Resource {
	var resources = make([]model.Resource, len(reservations))
	for i, reservation := range reservations {
		for _, instance := range reservation.Instances {
			resources[i] = model.Resource{
				CloudProvider: clouds.AWS,
				ResourceType:  "ec2",
				AccountId:     accountId,
				Region:        region,
				Name:          instance.InstanceId,
				CreationDate:  instance.LaunchTime,
				Tags:          awsToResourceTags(instance.Tags),
			}
		}
	}
	return resources
}

func getAwsRegions() []string {
	resolver := endpoints.DefaultResolver()
	partitions := resolver.(endpoints.EnumPartitions).Partitions()

	var regions []string
	for _, p := range partitions {
		for id := range p.Regions() {
			regions = append(regions, id)
		}
	}
	return regions
}

func awsToResourceTags(tags []*ec2.Tag) map[string]string {
	result := map[string]string{}
	for _, tag := range tags {
		result[*tag.Key] = *tag.Value
	}
	return result
}