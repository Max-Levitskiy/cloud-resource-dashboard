package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/conf"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/strategy"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"time"
)

var Client = elastic{
	es: getClient(),
}

func (e *elastic) SaveResource(resource model.Resource) {
	marshal, err := json.Marshal(resource)
	if err != nil {
		log.Panic(err)
	}
	resource.CloudId = resource.GenerateId()
	req := esapi.IndexRequest{
		Index:      conf.Inst.Elastic.Index.Resource.Name,
		DocumentID: resource.CloudId,
		Body:       bytes.NewReader(marshal),
		Refresh:    "true",
	}

	var res *esapi.Response
	err = retry.Retry(func(attempt uint) error {
		res, err = req.Do(context.Background(), e.es)
		if res != nil {
			defer res.Body.Close()
		}

		if err == nil {
			e.checkErrors(res)
			return nil
		} else {
			log.Println(err)
			return err
		}
	},
		strategy.Limit(10),
		strategy.Backoff(backoff.Fibonacci(10*time.Millisecond)),
	)
	if err != nil {
		log.Panic(err.Error())
	}
	e.checkErrors(res)
}

func (e *elastic) GetResourceById(documentId string) *model.Resource {
	request := esapi.GetRequest{
		Index:      conf.Inst.Elastic.Index.Resource.Name,
		DocumentID: documentId,
	}
	res, err := request.Do(context.Background(), e.es)
	if err == nil {
		defer res.Body.Close()
		e.checkErrors(res)
		var resource resourceEsResponse
		if err := json.NewDecoder(res.Body).Decode(&resource); err != nil {
			log.Panic(err)
		}
		resource.Source.CloudId = resource.Id
		return &resource.Source
	} else {
		log.Panic(err)
		return nil
	}
}

func (e *elastic) BulkSave(resources []model.Resource) {
	log.Println("Start bulk save")
	var (
		buf bytes.Buffer
		res *esapi.Response
		err error

		numItems  int
		currBatch int

		count int
		batch = 100
	)
	count = len(resources)

	for i, resource := range resources {
		numItems++
		resource.CloudId = resource.GenerateId()
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, resource.CloudId, "\n"))

		currBatch = i / batch
		if i == count-1 {
			currBatch++
		}
		data := append(e.toJson(resource), "\n"...)

		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)

		if i > 0 && i%batch == 0 || i == count-1 {
			if res, err = e.es.Bulk(bytes.NewReader(buf.Bytes()), e.es.Bulk.WithIndex(conf.Inst.Elastic.Index.Resource.Name)); err == nil {
				defer res.Body.Close()
				e.checkErrors(res)
			} else {
				log.Panic(err)
				return
			}
			buf.Reset()
			numItems = 0
		}
	}
	log.Println("Bulk save finished successfully")

}

func (e *elastic) toJson(resource model.Resource) []byte {
	marshaled, err := json.Marshal(resource)
	if err != nil {
		log.Panic(err)
	}
	return marshaled
}

func (e *elastic) ClearResourceIndex() {
	log.Println("Deleting resource index...")
	request := esapi.IndicesDeleteRequest{
		Index: []string{conf.Inst.Elastic.Index.Resource.Name},
	}
	_, err := request.Do(context.Background(), e.es)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Done")
}

func getClient() *elasticsearch.Client {
	log.Println("Create new ES client")
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:30920",
		},
	})
	if err == nil {
		return es
	} else {
		log.Panic(err)
		return nil
	}
}

func (e *elastic) checkErrors(res *esapi.Response) {
	if res.IsError() {
		var body map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			log.Panicf("Error parsing the response body: %s", err)
		} else {
			log.Println(body)
			if err, ok := body["error"]; ok {
				log.Panicf("[%s] %s", res.Status(), err)
			} else if found, ok := body["found"]; ok {
				if found == false {
					log.Panicf("Document %s not found in index %s", body["_id"], body["_index"])
				} else {
					log.Panicf("Unknown error. %s", body)
				}
			}
		}
	}
}

func (e *elastic) CreateIndex() {
	req := esapi.IndicesCreateRequest{
		Index: conf.Inst.Elastic.Index.Resource.Name,
	}
	if res, err := req.Do(context.Background(), e.es); err != nil {
		log.Panic(err)
	} else {
		e.checkErrors(res)
		e.updateIndexMapping()
	}
}

func (e *elastic) updateIndexMapping() {
	req := esapi.IndicesPutMappingRequest{
		Index: []string{conf.Inst.Elastic.Index.Resource.Name},
		Body: bytes.NewBuffer([]byte(`
		{
			"properties": {
				"Service": {
					"type": "keyword"
				},
				"Region": {
					"type": "keyword"
				},
				"CloudProvider": {
					"type": "keyword"
				},
				"AccountId": {
					"type": "keyword"
				},
				"ResourceId": {
					"type": "keyword"
				},
				"Tags": {
					"type": "flattened"
				},
				"CreationDate": {
					"type": "date"
				}
			}
		}
		`)),
	}
	if res, err := req.Do(context.Background(), e.es); err != nil {
		log.Panic(err)
	} else {
		e.checkErrors(res)
	}
}
