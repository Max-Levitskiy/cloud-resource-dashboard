package conf

import (
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/reflection"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	tests := []func(t *testing.T){
		shouldInitFieldsByDefault,
		shouldAddConfig,
		shouldOverrideDefaultFromFile,
		shouldResetConfig,
		shouldOverrideConfigFromEnv,
		shouldOverrideFilesByEnv,
		shouldReadAwsProfiles,
		shouldReadHomeDir,
		shouldReadHomeDirFromEnv,
		shouldUseConfPathVariable,
	}

	for _, test := range tests {
		t.Run(reflection.GetFunctionName(test), test)
		Reset()
	}
}

func shouldUseConfPathVariable(t *testing.T) {
	defer unsetEnv(t, "CONFIG_PATH")
	setEnv(t, "CONFIG_PATH", "/tmp/")
	defer os.Remove("/tmp/config.yaml")
	err := ioutil.WriteFile("/tmp/config.yaml", []byte(`
elastic:
  server: "tstServ"
`), os.ModePerm)
	assert.Nil(t, err)

	Reset()

	assert.Equal(t, Inst.Elastic.Server, "tstServ")
}

func shouldReadHomeDir(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	assert.Nil(t, err)

	assert.Equal(t, homeDir, Inst.HomeDir)
}

func shouldReadHomeDirFromEnv(t *testing.T) {
	homeDir := "/some/home/dir"
	setEnv(t, "HOME_DIR", homeDir)
	Reset()
	assert.Equal(t, homeDir, Inst.HomeDir)
}

func shouldInitFieldsByDefault(t *testing.T) {
	checkDefaultConfigs(t)
}

func shouldAddConfig(t *testing.T) {
	AddConfigs("test")

	assert.Equal(t, "resources_test", Inst.Elastic.Index.Resource.Name)
}

func shouldOverrideDefaultFromFile(t *testing.T) {
	defer unsetEnv(t, "CONFIG_FILE_POSTFIX")
	setEnv(t, "CONFIG_FILE_POSTFIX", "test")

	Reset()
	assert.Equal(t, "resources_test", Inst.Elastic.Index.Resource.Name)
	assert.Equal(t, 30920, Inst.Elastic.Port)
}

func shouldResetConfig(t *testing.T) {
	AddConfigs("test")
	Reset()

	checkDefaultConfigs(t)
}

func shouldOverrideConfigFromEnv(t *testing.T) {
	serverName := "someServer"
	defer unsetEnv(t, "ELASTIC_SERVER")
	setEnv(t, "ELASTIC_SERVER", serverName)
	Reset()

	assert.Equal(t, serverName, Inst.Elastic.Server)
}

func shouldOverrideFilesByEnv(t *testing.T) {
	indexName := "someName"
	setEnv(t, "ELASTIC_INDEX_RESOURCE_NAME", indexName)
	AddConfigs("test")

	assert.Equal(t, indexName, Inst.Elastic.Index.Resource.Name)
}

func shouldReadAwsProfiles(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.Nil(t, err)
	credentialsFilePath := currentDir + "/credentials"
	setEnv(t, "AWS_CONFIG_PATH", currentDir)
	defer os.Remove(credentialsFilePath)
	p1 := "profile1"
	p2 := "profile2"
	assert.Nil(t, ioutil.WriteFile(
		credentialsFilePath,
		[]byte(fmt.Sprintf("[%s]\nsomecontent\n[%s]\nmorecontent", p1, p2)),
		os.ModePerm,
	))

	Reset()

	assert.Len(t, Inst.AWS.ProfileNames, 2)
	assert.Contains(t, Inst.AWS.ProfileNames, &p1, &p2)
}

func setEnv(t *testing.T, key string, value string) {
	assert.Nil(t, os.Setenv(key, value))
}

func unsetEnv(t *testing.T, key string) {
	assert.Nil(t, os.Unsetenv(key))
}

func checkDefaultConfigs(t *testing.T) {
	assert.Equal(t, "localhost", Inst.Elastic.Server)
	assert.Equal(t, 30920, Inst.Elastic.Port)
	assert.Equal(t, "resources", Inst.Elastic.Index.Resource.Name)
}
