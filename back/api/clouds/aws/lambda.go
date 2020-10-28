package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/sirupsen/logrus"
)

func scanLambdaFunctions(accountId *string, region string) {
	logrus.Infof("Start scan LambdaFunctions for %s account %s region", *accountId, region)
	lambdaList, err := ListEBS(region)
	if err == nil {
		resources := LambdaInstancesToResources(lambdaList.Volumes, accountId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	logrus.Infof("Scan LambdaFunctions for %s region finished", region)
}

func ListLambdaFunctions(region string) (*lambda.ListFunctionsOutput, error) {
	input := &lambda.ListFunctionsInput{}
	session := getSession(region)

	svc := lambda.New(session)

	result, err := svc.ListFunctions(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func LambdaInstancesToResources(volumes []*ec2.Volume, accountId *string, region *string) []*model.Resource {
	var resources = make([]*model.Resource, len(volumes))
	for i, volume := range volumes {

		resources[i] = &model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "lambda",
			AccountId:     accountId,
			Region:        region,
			ResourceId:    volume.VolumeId,
			CreationDate:  volume.CreateTime,
			Tags:          awsToResourceTags(volume.Tags),
		}
	}
	return resources
}
