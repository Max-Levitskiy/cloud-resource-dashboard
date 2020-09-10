package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/strategy"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"time"
)

const (
	ResourcesIndex = "resources"
)

var es = getClient()

func SaveResource(resource model.Resource) {
	marshal, err := json.Marshal(resource)
	if err != nil {
		log.Panic(err)
	}
	req := esapi.IndexRequest{
		Index:      ResourcesIndex,
		DocumentID: resource.GenerateId(),
		Body:       bytes.NewReader(marshal),
		Refresh:    "true",
	}

	var res *esapi.Response
	err = retry.Retry(func(attempt uint) error {
		defer res.Body.Close()
		res, err = req.Do(context.Background(), es)

		if err == nil && !res.IsError() {
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

	if res.IsError() {
		log.Panic(res.StatusCode, res)
	}
}

func BulkSave(resources []model.Resource) {
	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: getClient(),
	})

	for _, resource := range resources {
		err := indexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   toJson(resource),
			})
		if err != nil {
			log.Panic(err)
		}
	}

	err := indexer.Close(context.Background())
	if err != nil {
		log.Panic(err.Error())
	}
	//
	//var requestBody []byte
	//requestBody= append(requestBody, "\n"...)
	//
	//req := esapi.BulkRequest{
	//	Index:      ResourcesIndex,
	//	Body:       bytes.NewReader(requestBody),
	//	Refresh:    "true",
	//}
	//
	//es.Bulk.WithDocumentType("")
	//var err error
	//err = retry.Retry(func(attempt uint) error {
	//
	//	res, err := req.Do(context.Background(), es)
	//	if err != nil {
	//		log.Println(err)
	//	} else if res.IsError() {
	//		defer res.Body.Close()
	//		buf := new(bytes.Buffer)
	//		_, _ = buf.ReadFrom(res.Body)
	//		log.Panic(buf.String())
	//	}
	//	return err
	//},
	//	strategy.Limit(10),
	//	strategy.Backoff(backoff.Fibonacci(10*time.Millisecond)),
	//)
	//if err != nil {
	//	log.Panic(err.Error())
	//}
}

func toJson(resource model.Resource) *bytes.Reader {
	marshaled, err := json.Marshal(resource)
	if err != nil {
		log.Panic(err)
	}
	return bytes.NewReader(marshaled)
}

func ClearResourceIndex() {
	log.Println("Deleting resource index...")
	maxDoc := 2
	request := esapi.DeleteByQueryRequest{
		Index:   []string{ResourcesIndex},
		MaxDocs: &maxDoc,
	}
	_, err := request.Do(context.Background(), es)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Done")
}

func getClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://localhost:30920",
		},
	})
	if err == nil {
		return es
	} else {
		log.Panic(err)
		return nil
	}
}
