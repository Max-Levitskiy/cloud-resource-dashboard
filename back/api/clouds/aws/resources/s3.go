package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Scanner struct{}

var unknownRegion = "unknown"

func (S3Scanner) Scan(projectId *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error) {
	s := session.GetForDefaultRegion(profileName)
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

func listS3(s *awsSession.Session) (*s3.ListBucketsOutput, error) {
	svc := s3.New(s)
	return svc.ListBuckets(&s3.ListBucketsInput{})
}

func fetchRegion(bucketName *string, s *awsSession.Session) *string {
	input := &s3.GetBucketLocationInput{
		Bucket: bucketName,
	}
	svc := s3.New(s)

	bucketLocation, err := svc.GetBucketLocation(input)
	if err == nil && bucketLocation.LocationConstraint != nil {
		return bucketLocation.LocationConstraint
	} else {
		logger.Warn.Print("Can't fetch region for bucket ", *bucketName, "Error: ", err)
		return &unknownRegion
	}
}
