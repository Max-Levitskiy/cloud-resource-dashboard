package conf

import (
	"bufio"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/go-homedir"
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
	AWS struct {
		ProfileNames []*string
	}
}

var Inst = initConfig()

func initConfig() config {
	var cfg = config{}
	confPostfix := os.Getenv("CONFIG_FILE_POSTFIX")
	addConfigs(&cfg, "", confPostfix)
	readAwsProfiles(&cfg)
	return cfg
}

func readAwsProfiles(cfg *config) {
	var configPath string
	if configPath = os.Getenv("AWS_CONFIG_PATH"); configPath == "" {
		if homeDir, err := homedir.Dir(); err != nil {
			logrus.Error(err)
		} else {
			configPath = homeDir + "/.aws"
		}
	}
	if file, err := os.Open(configPath + "/credentials"); err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				line = strings.TrimPrefix(line, "[")
				line = strings.TrimSuffix(line, "]")
				cfg.AWS.ProfileNames = append(cfg.AWS.ProfileNames, &line)
			}
		}
	} else {
		logrus.Error(err)
	}
}

func AddConfigs(filesPostfixes ...string) {
	addConfigs(&Inst, filesPostfixes...)
}
func addConfigs(config *config, filesPostfixes ...string) {
	for _, postfix := range filesPostfixes {
		readFile(config, postfix)
	}
	readEnv(config)
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
