package model

import (
	"strings"
	"time"
)

type Resource struct {
	CloudProvider string
	ResourceType  string
	AccountId     *string
	Name          *string
	CreationDate  *time.Time
	Tags          []string
}

func (r *Resource) GenerateId() string {
	return strings.Join([]string{r.CloudProvider, r.ResourceType, *r.AccountId}, "-")
}
