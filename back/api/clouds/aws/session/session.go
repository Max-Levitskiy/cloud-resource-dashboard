package session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"os"
	"sync"
)

var activeSessionsCache sync.Map
var noProfile = "[noprofile]"

const DefaultRegion = "us-east-1"

func GetForDefaultRegion(profile ...*string) *session.Session {
	return Get(DefaultRegion, profile...)
}
func Get(region string, profile ...*string) *session.Session {
	if len(profile) > 1 {
		log.Panic("You can use no more then 1 profile. Given: ", profile)
	}
	profileGiven := len(profile) == 1
	var profileName *string
	if profileName = &noProfile; profileGiven {
		profileName = profile[0]
	}
	cacheKey := fmt.Sprintf("%s:%s", *profileName, region)
	if s, ok := activeSessionsCache.Load(cacheKey); ok {
		return s.(*session.Session)
	} else {
		options := session.Options{
			Config: aws.Config{
				Region:                        &region,
				CredentialsChainVerboseErrors: aws.Bool(true),
			},
		}
		if len(profile) == 1 {
			_ = os.Unsetenv("AWS_ACCESS_KEY_ID")
			_ = os.Unsetenv("AWS_SECRET_KEY")
			_ = os.Unsetenv("AWS_SECRET_ACCESS_KEY")
			options.Profile = *profile[0]
		}
		s, err := session.NewSessionWithOptions(options)
		if err == nil {
			activeSessionsCache.Store(cacheKey, s)
			return s
		} else {
			log.Panic("Can't get session for region", region, err)
			return nil
		}
	}
}
