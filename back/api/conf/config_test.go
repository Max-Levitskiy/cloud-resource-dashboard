package conf

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/reflection"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	tests := []func(t *testing.T){
		shouldInitFieldsByDefault,
		shouldAddConfig,
		shouldResetConfig,
		shouldOverrideConfigFromEnv,
		shouldOverrideFilesByEnv,
	}

	for _, test := range tests {
		t.Run(reflection.GetFunctionName(test), test)
	}
}

func shouldInitFieldsByDefault(t *testing.T) {
	checkDefaultConfigs(t)
}

func shouldAddConfig(t *testing.T) {
	AddConfigs("test")

	assert.Equal(t, "resources_test", Inst.Elastic.Index.Resource.Name)
}

func shouldResetConfig(t *testing.T) {
	AddConfigs("test")
	Reset()

	checkDefaultConfigs(t)
}

func shouldOverrideConfigFromEnv(t *testing.T) {
	serverName := "someServer"
	assert.Nil(t, os.Setenv("ELASTIC_SERVER", serverName))
	Reset()

	assert.Equal(t, serverName, Inst.Elastic.Server)
}

func shouldOverrideFilesByEnv(t *testing.T) {
	indexName := "someName"
	assert.Nil(t, os.Setenv("ELASTIC_INDEX_RESOURCE_NAME", indexName))
	AddConfigs("test")

	assert.Equal(t, indexName, Inst.Elastic.Index.Resource.Name)
}

func checkDefaultConfigs(t *testing.T) {
	assert.Equal(t, "localhost", Inst.Elastic.Server)
	assert.Equal(t, 30920, Inst.Elastic.Port)
	assert.Equal(t, "resources", Inst.Elastic.Index.Resource.Name)
}
