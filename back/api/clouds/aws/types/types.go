package types

import "github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"

type RegionalResourceScanner interface {
	Scan(projectId *string, region *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error)
}
type GlobalResourceScanner interface {
	Scan(projectId *string, profileName *string, saveCh chan<- *model.Resource, errCh chan<- error)
}
