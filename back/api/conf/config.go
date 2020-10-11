package conf

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

//
type config struct {
	Elastic struct {
		Server string `yaml:"server" envconfig:"ELASTIC_SERVER"`
		Port   int    `yaml:"port" envconfig:"ELASTIC_PORT"`
		Index  struct {
			Resource struct {
				Name string `yaml:"name" envconfig:"ELASTIC_INDEX_RESOURCE_NAME"`
			} `yaml:"resource"`
		} `yaml:"index"`
	} `yaml:"elastic"`
}

var Inst = initConfig()

func initConfig() config {
	var cfg = config{}
	confPostfix := os.Getenv("CONFIG_FILE_POSTFIX")
	readFile(&cfg, confPostfix)
	readEnv(&cfg)
	return cfg
}

func AddConfigs(filesPostfixes ...string) {
	for _, postfix := range filesPostfixes {
		readFile(&Inst, postfix)
	}
	readEnv(&Inst)
}

func Reset() {
	Inst = initConfig()
}

func readFile(cfg *config, filePostfix string) {
	filePostfix = strings.TrimSpace(filePostfix)
	var fileName string
	if filePostfix == "" {
		fileName = "config.yaml"
	} else {
		fileName = fmt.Sprintf("config_%s.yaml", filePostfix)
	}

	f, err := os.Open("conf/" + fileName)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
func processError(err error) {
	fmt.Println(err)
	logrus.Fatal("Can't read config. Exit.")
}
