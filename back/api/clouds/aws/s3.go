package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func scanS3(accountId *string, region string) {
	log.Printf("Start scan S3 for %s region", region)
	listS3, err := ListS3(region)
	if err == nil {
		resources := s3BucketsToResources(listS3.Buckets, accountId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	log.Printf("Scan S3 for %s region finished", region)
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

func s3BucketsToResources(buckets []*s3.Bucket, accountId *string, region *string) []model.Resource {
	var resources = make([]model.Resource, len(buckets))
	for i, bucket := range buckets {
		resources[i] = model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "s3",
			AccountId:     accountId,
			Region:        region,
			ResourceId:    bucket.Name,
			CreationDate:  bucket.CreationDate,
		}
	}
	return resources
}
