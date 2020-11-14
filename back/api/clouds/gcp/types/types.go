package types

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp/session"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
)

type Scanner interface {
	Scan(s *session.Session, saveCh chan<- *model.Resource, errCh chan<- error)
}
