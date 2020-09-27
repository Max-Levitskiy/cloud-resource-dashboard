package model

import (
	"strings"
	"time"
)

type Resource struct {
	CloudId       string
	CloudProvider string
	Service       string
	AccountId     *string
	ResourceId    *string
	Region        *string
	CreationDate  *time.Time
	Tags          map[string]string
}

func (r *Resource) GenerateId() string {
	return strings.Join([]string{r.CloudProvider, r.Service, *r.Region, *r.AccountId, *r.ResourceId}, ":")
}
