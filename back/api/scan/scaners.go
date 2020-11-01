package scan

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/aws"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/clouds/gcp"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/persistanse/elasticsearch"
)

var providers = map[string]func(){
	clouds.AWS: aws.FullScan,
	clouds.GCP: gcp.FullScan,
}

func StartFullScan() {
	elasticsearch.Client.ClearResourceIndex()
	elasticsearch.Client.CreateIndex()
	for _, f := range providers {
		go f()
	}
}
