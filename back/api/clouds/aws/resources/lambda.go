package resources

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaScanner struct{}

func (l LambdaScanner) Scan(projectId *string, region *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error) {
	lambdaList, err := listLambdaFunctions(region, profileName)
	if err == nil {
		for _, function := range lambdaList.Functions {
			saveCh <- &model.Resource{
				CloudProvider: clouds.AWS,
				Service:       "lambda",
				ProjectId:     projectId,
				Region:        region,
				ResourceId:    function.FunctionArn,
			}
		}
	} else {
		errCh <- err
	}
}

func listLambdaFunctions(region *string, profileName *string) (*lambda.ListFunctionsOutput, error) {
	svc := lambda.New(session.Get(region, profileName))
	return svc.ListFunctions(&lambda.ListFunctionsInput{})
}
