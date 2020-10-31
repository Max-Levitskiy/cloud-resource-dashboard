package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	session2 "github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/sirupsen/logrus"
)

func ScanLambdaFunctions(accountId *string, region string) {
	logrus.Infof("Start scan LambdaFunctions for %s account %s region", *accountId, region)
	lambdaList, err := listLambdaFunctions(region)
	if err == nil {
		resources := LambdaInstancesToResources(lambdaList.Functions, accountId, &region)
		elasticsearch.Client.BulkSave(resources)
	}
	logrus.Infof("Scan LambdaFunctions for %s region finished", region)
}

func listLambdaFunctions(region string) (*lambda.ListFunctionsOutput, error) {
	input := &lambda.ListFunctionsInput{}
	session := session2.Get(region)

	svc := lambda.New(session)

	result, err := svc.ListFunctions(input)
	if err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func LambdaInstancesToResources(lambdas []*lambda.FunctionConfiguration, accountId *string, region *string) []*model.Resource {
	var resources = make([]*model.Resource, len(lambdas))
	for i, lambda := range lambdas {

		resources[i] = &model.Resource{
			CloudProvider: clouds.AWS,
			Service:       "lambda",
			AccountId:     accountId,
			Region:        region,
			ResourceId:    lambda.FunctionArn,
		}
	}
	return resources
}
