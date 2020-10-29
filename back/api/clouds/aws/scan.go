package aws

import (
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/conf"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"reflect"
)

var awsRegions = getAwsRegions()

type regionalResourceScanner interface {
	scan(accountId *string, region string, profileName string)
}
type globalResourceScanner interface {
	scan(accountId *string, profileName *string)
}

const defaultRegion = "us-east-1"

var regionGlobalScanners = []globalResourceScanner{
	s3Scanner{},
}

func FullScan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	logger.Info.Println("Start AWS scan")
	for _, profileName := range conf.Inst.AWS.ProfileNames {
		accountId := getAccountId(profileName)
		for _, scanner := range regionGlobalScanners {
			go func(s globalResourceScanner, accountId *string, profileName *string) {
				scannerType := reflect.TypeOf(s).String()
				logger.Info.Printf("Starting %s scan for profile %s, account id: %s", scannerType, *profileName, *accountId)
				s.scan(accountId, profileName)
				logger.Info.Printf("Finished %s for profile %s, account id: %s", scannerType, *profileName, *accountId)
			}(scanner, accountId, profileName)
		}
		for _, region := range awsRegions {
			//go scanS3(accountId, region, profileName)
			go scanEc2(accountId, region)
			go scanEBS(accountId, region)
			go scanElb(accountId, region)
			go scanLambdaFunctions(accountId, region)
			go scanVpc(accountId, region)
		}
	}
}

func getAccountId(profileName *string) *string {
	s := sts.New(getSession(defaultRegion, profileName))
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

func awsToResourceTags(tags []*ec2.Tag) map[string]string {
	result := map[string]string{}
	for _, tag := range tags {
		result[*tag.Key] = *tag.Value
	}
	return result
}
