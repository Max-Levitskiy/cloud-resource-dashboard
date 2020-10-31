package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	session2 "github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/types"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"sync"
)

type S3Scanner struct {
	types.GlobalResourceScanner
}

var unknownRegion = "unknown"

func (S3Scanner) Scan(accountId *string, profileName *string) {
	s := session2.GetForDefaultRegion(profileName)
	listS3, err := listS3(s)
	if err == nil {
		resources := s3BucketsToResources(listS3.Buckets, accountId, s)
		elasticsearch.Client.BulkSave(resources)
	}
	log.Printf("Scan S3 for profile %s finished", *profileName)
}

func listS3(s *session.Session) (*s3.ListBucketsOutput, error) {
	input := &s3.ListBucketsInput{}
	svc := s3.New(s)

	result, err := svc.ListBuckets(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func s3BucketsToResources(buckets []*s3.Bucket, accountId *string, s *session.Session) []*model.Resource {
	var resources = make([]*model.Resource, len(buckets))
	for i, bucket := range buckets {
		resources[i] = &model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "s3",
			AccountId:     accountId,
			ResourceId:    bucket.Name,
			CreationDate:  bucket.CreationDate,
		}
	}
	fetchBucketsRegions(&resources, s)

	return resources
}

func fetchBucketsRegions(buckets *[]*model.Resource, s *session.Session) {
	var wg sync.WaitGroup
	wg.Add(len(*buckets))
	for _, bucket := range *buckets {
		go func(b *model.Resource, wg *sync.WaitGroup) {
			b.Region = fetchRegion(b.ResourceId, s)
			wg.Done()
		}(bucket, &wg)
	}
	wg.Wait()
}

func fetchRegion(bucketName *string, s *session.Session) *string {
	input := &s3.GetBucketLocationInput{
		Bucket: bucketName,
	}
	svc := s3.New(s)

	bucketLocation, err := svc.GetBucketLocation(input)
	if err == nil && bucketLocation.LocationConstraint != nil {
		return bucketLocation.LocationConstraint
	} else {
		log.Print("Can't fetch region for bucket ", *bucketName, "Error: ", err)
		return &unknownRegion
	}
}
