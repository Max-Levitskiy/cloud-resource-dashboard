package aws_model

import "github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"

type AwsResource interface {
	model.IResource
	GetArn() *string
}

type S3Resource struct {
	AwsResource
}
