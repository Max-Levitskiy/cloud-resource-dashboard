package model

import (
	"strings"
	"time"
)

type IResource interface {
	getId() string
	getCloudProvider() string
	getResourceType() string
	getAccountId() *string
	getName() *string
	getRegion() *string
	getCreationDate() *time.Time
	getTags() map[string]string
}

type Resource struct {
	Id            string
	CloudProvider string
	ResourceType  string
	AccountId     *string
	Name          *string
	Region        *string
	CreationDate  *time.Time
	Tags          map[string]string
}

func (r *Resource) GenerateId() string {
	return strings.Join([]string{r.CloudProvider, r.ResourceType, *r.Region, *r.AccountId, *r.Name}, "-")
}
