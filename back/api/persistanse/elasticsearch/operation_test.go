package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/conf"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/model"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestSaveResource(t *testing.T) {
	defer Client.ClearResourceIndex()

	resource := getResourceExample()
	Client.SaveResource(resource)

	resource.Id = resource.GenerateId()
	savedResource := Client.GetResourceById(resource.Id)

	assert.Equal(t, resource.Id, savedResource.Id)
	assert.Equal(t, resource.ResourceType, savedResource.ResourceType)
	assert.Equal(t, resource.AccountId, savedResource.AccountId)
	assert.Equal(t, resource.Name, savedResource.Name)
	assert.Equal(t, resource.Tags, savedResource.Tags)
	assert.True(t, resource.CreationDate.Equal(*savedResource.CreationDate))
}

func TestBulkSave(t *testing.T) {
	defer Client.ClearResourceIndex()

	resource := getResourceExample()

	Client.BulkSave([]model.Resource{resource})

	resource.Id = resource.GenerateId()
	savedResource := Client.GetResourceById(resource.Id)

	assert.Equal(t, resource.Id, savedResource.Id)
	assert.Equal(t, resource.ResourceType, savedResource.ResourceType)
	assert.Equal(t, resource.AccountId, savedResource.AccountId)
	assert.Equal(t, resource.Name, savedResource.Name)
	assert.Equal(t, resource.Tags, savedResource.Tags)
	assert.True(t, resource.CreationDate.Equal(*savedResource.CreationDate))
}

func TestCreateIndex(t *testing.T) {
	defer Client.ClearResourceIndex()

	Client.CreateIndex()

	req := esapi.IndicesExistsRequest{
		Index: []string{conf.Inst.Elastic.Index.Resource.Name},
	}

	if response, err := req.Do(context.Background(), getClient()); err != nil {
		log.Panic(err)
	} else {
		log.Println(response)
		assert.False(t, response.IsError())
		assert.Equal(t, 200, response.StatusCode)
	}
}

func TestUpdateIndexMapping(t *testing.T) {
	// tear down
	defer Client.ClearResourceIndex()

	// given
	Client.CreateIndex()

	// when
	Client.updateIndexMapping()

	// then
	req := esapi.IndicesGetMappingRequest{
		Index: []string{conf.Inst.Elastic.Index.Resource.Name},
	}

	res, err := req.Do(context.Background(), getClient())

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	var body map[string]map[string]map[string]map[string]map[string]string
	assert.Nil(t, json.NewDecoder(res.Body).Decode(&body))

	assert.Equal(t, "keyword", body["resources"]["mappings"]["properties"]["Name"]["type"])
}

func getResourceExample() model.Resource {
	accountId := "someAccId"
	name := "aName"
	region := "us-east-1"
	date := time.Now()
	resource := model.Resource{
		CloudProvider: "AWS",
		ResourceType:  "testType",
		AccountId:     &accountId,
		Name:          &name,
		Region:        &region,
		CreationDate:  &date,
		Tags:          map[string]string{"a": "b"},
	}
	return resource
}
