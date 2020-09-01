package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var activeSessionsCache = make(map[string]*session.Session)

func getSession(region string) (*session.Session, error) {
	if s, ok := activeSessionsCache[region]; ok {
		return s, nil
	} else {
		s, err := session.NewSession(&aws.Config{
			Region: &region,
		})
		if err == nil {
			return s, nil
		} else {
			return nil, err
		}
	}

}
