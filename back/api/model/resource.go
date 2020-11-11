package model

import (
	"strings"
	"time"
)

type Resource struct {
	CloudId       string
	CloudProvider string
	Service       string
	ProjectId     *string
	ResourceId    *string
	Region        *string
	CreationDate  *time.Time
	Tags          map[string]string
}

func (r *Resource) GenerateId() string {
	if r.CloudId == "" {
		return strings.Join([]string{r.CloudProvider, r.Service, *r.Region, *r.ProjectId, *r.ResourceId}, ":")
	} else {
		return r.CloudId
	}
}
