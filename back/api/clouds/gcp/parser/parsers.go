package parser

import (
	"regexp"
	"strings"
)

var reg = compileNameRegexp()

type GcpResourceName struct {
	ProjectId    string
	Location     string
	ResourceType string
	ResourceName string
}

func ParseName(name string) GcpResourceName {
	parsed := reg.FindAllStringSubmatch(name, -1)[0]
	return GcpResourceName{
		ProjectId:    parsed[1],
		Location:     parsed[2],
		ResourceType: parsed[3],
		ResourceName: parsed[4],
	}
}
func GetCleanRegionName(fullRegion string) string {
	regionSplit := strings.Split(fullRegion, "/")
	return regionSplit[len(regionSplit)-1]
}

func compileNameRegexp() *regexp.Regexp {
	return regexp.MustCompile(`projects/([a-zA-Z0-9-]+)/locations/([a-zA-Z0-9-]+)/([a-zA-Z0-9-]+)/([a-zA-Z0-9-]+)`)
}
