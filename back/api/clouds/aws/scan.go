package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
)

var awsRegions = getAwsRegions()

func FullScan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	log.Println("Start AWS scan")
	accountId := getAccountId()
	for _, region := range awsRegions {
		go scanS3(accountId, region)
		go scanEc2(accountId, region)
		go scanEBS(accountId, region)
		go scanElb(accountId, region)
		go scanLambdaFunctions(accountId, region)
		go scanVpc(accountId, region)
	}
}

func getAccountId() *string {
	s := sts.New(getSession("us-east-1"))
	identity, err := s.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err == nil {
		return identity.Account
	} else {
		log.Panic(err)
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
