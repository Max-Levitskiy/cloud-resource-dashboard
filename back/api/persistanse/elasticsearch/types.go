package elasticsearch

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/elastic/go-elasticsearch/v8"
)

type resourceEsResponse struct {
	Id     string         `json:"_id"`
	Source model.Resource `json:"_source"`
}
type elastic struct {
	es *elasticsearch.Client
}
