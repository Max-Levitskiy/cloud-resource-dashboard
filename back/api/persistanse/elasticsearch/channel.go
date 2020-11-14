package elasticsearch

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"time"
)

var BulkSaveCh = initChan()

func initChan() chan *model.Resource {
	ch := make(chan *model.Resource)
	var resources []*model.Resource
	go func() {
		timer := time.Now()
		for {
			select {
			case r := <-ch:
				resources = append(resources, r)
			case <-time.After(3 * time.Second):
			}
			resCount := len(resources)
			if resCount > 100 || time.Now().Sub(timer) > 3*time.Second {
				if resCount > 0 {
					Client.BulkSave(resources)
					resources = nil
				}
				timer = time.Now()
			}
		}
	}()
	return ch
}
