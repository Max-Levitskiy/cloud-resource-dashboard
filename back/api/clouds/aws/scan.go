package aws

import (
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/resources"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/types"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/conf"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/sts"
	"reflect"
)

var (
	awsRegions     = getAwsRegions()
	globalScanners = []types.GlobalResourceScanner{
		resources.S3Scanner{},
	}
	regionScanners = []types.RegionalResourceScanner{
		resources.EbsScanner{},
		resources.Ec2Scanner{},
		resources.VpcScanner{},
		resources.ElbScanner{},
		resources.LambdaScanner{},
	}
)

func FullScan(saveCh chan<- *model.Resource, errCh chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Warn.Println("Recovered: ", r)
		}
	}()
	logger.Info.Println("Start AWS Scan")
	for _, profileName := range conf.Inst.AWS.ProfileNames {
		projectId, err := getProjectId(profileName)
		if err != nil {
			errCh <- err
			continue
		}

		for _, region := range awsRegions {
			for _, scanner := range regionScanners {
				go func(region string, s types.RegionalResourceScanner, projectId string, profileName string) {
					scannerType := reflect.TypeOf(s).String()
					logger.Info.Printf("Starting %s Scan for profile %s, account id: %s", scannerType, profileName, projectId)
					s.Scan(&projectId, &region, &profileName, saveCh, errCh)
					logger.Info.Printf("Finished %s for profile %s, account id: %s", scannerType, profileName, projectId)
				}(region, scanner, projectId, profileName)
			}
		}
		for _, scanner := range globalScanners {
			go func(s types.GlobalResourceScanner, projectId string, profileName string) {
				scannerType := reflect.TypeOf(s).String()
				logger.Info.Printf("Starting %s Scan for profile %s, account id: %s", scannerType, profileName, projectId)
				s.Scan(&projectId, &profileName, saveCh, errCh)
				logger.Info.Printf("Finished %s for profile %s, account id: %s", scannerType, profileName, projectId)
			}(scanner, projectId, profileName)
		}
	}
}

func getProjectId(profileName string) (string, error) {
	s := sts.New(session.Get(&session.DefaultRegion, &profileName))
	identity, err := s.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err == nil {
		return *identity.Account, nil
	} else {
		return "", fmt.Errorf("can't get project id by profile name %s. Error: %v", profileName, err)
	}
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
