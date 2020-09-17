package model

import (
	"strings"
	"time"
)

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
