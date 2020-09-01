package aws

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListS3(region string) (*s3.ListBucketsOutput, error) {
	input := &s3.ListBucketsInput{}
	session, err := getSession(region)
	if err != nil {
		return nil, err
	}

	svc := s3.New(session)

	result, err := svc.ListBuckets(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}
