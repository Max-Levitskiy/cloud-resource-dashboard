package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

var activeSessionsCache = make(map[string]*session.Session)

func getSession(region string) *session.Session {
	if s, ok := activeSessionsCache[region]; ok {
		return s
	} else {
		s, err := session.NewSession(&aws.Config{
			Region: &region,
		})
		if err == nil {
			return s
		} else {
			log.Panic("Can't get session for region", region, err)
			return nil
		}
	}
}
