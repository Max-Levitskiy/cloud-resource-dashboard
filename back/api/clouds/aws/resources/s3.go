package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	session2 "github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/types"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type S3Scanner struct {
	types.GlobalResourceScanner
}

var unknownRegion = "unknown"

func (S3Scanner) Scan(projectId *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error) {
	s := session2.GetForDefaultRegion(profileName)
	listS3, err := listS3(s)
	if err == nil {
		for _, bucket := range listS3.Buckets {
			go func(b *s3.Bucket) {
				r := &model.Resource{
					CloudProvider: clouds.AWS,
					Service:       "s3",
					ProjectId:     projectId,
					ResourceId:    b.Name,
					CreationDate:  b.CreationDate,
				}
				r.Region = fetchRegion(b.Name, s)
				saveCh <- r
			}(bucket)
		}
	} else {
		errCh <- err
	}
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
