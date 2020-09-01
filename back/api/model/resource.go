package model

import "time"

type Resource struct {
	CloudProvider string
	Name          *string
	CreationDate  *time.Time
	Tags          []string
}
