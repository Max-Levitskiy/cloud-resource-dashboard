package aws

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/resources"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws/types"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/conf"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/sts"
	"reflect"
)

var awsRegions = getAwsRegions()

var regionGlobalScanners = []types.GlobalResourceScanner{
	resources.S3Scanner{},
}

func FullScan() {
	defer func() {
		if r := recover(); r != nil {
			logger.Warn.Println("Recovered: ", r)
		}
	}()
	logger.Info.Println("Start AWS Scan")
	for _, profileName := range conf.Inst.AWS.ProfileNames {
		accountId := getAccountId(profileName)
		for _, scanner := range regionGlobalScanners {
			go func(s types.GlobalResourceScanner, accountId *string, profileName *string) {
				scannerType := reflect.TypeOf(s).String()
				logger.Info.Printf("Starting %s Scan for profile %s, account id: %s", scannerType, *profileName, *accountId)
				s.Scan(accountId, profileName)
				logger.Info.Printf("Finished %s for profile %s, account id: %s", scannerType, *profileName, *accountId)
			}(scanner, accountId, profileName)
		}
		for _, region := range awsRegions {
			//go scanS3(accountId, region, profileName)
			go resources.ScanEc2(accountId, region)
			go resources.ScanEBS(accountId, region)
			go resources.ScanElb(accountId, region)
			go resources.ScanLambdaFunctions(accountId, region)
			go resources.ScanVpc(accountId, region)
		}
	}
}

func getAccountId(profileName *string) *string {
	s := sts.New(session.Get(session.DefaultRegion, profileName))
	identity, err := s.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err == nil {
		return identity.Account
	} else {
		logger.Error.Panic(err)
		return nil
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
