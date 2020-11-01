package conf

import (
	"bufio"
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"runtime"
	"strings"
)

type config struct {
	HomeDir string `envconfig:"HOME_DIR"`
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
		ConfigPath   string
		ProfileNames []*string
	}
	GCP struct {
		CredentialsPath *string `yaml:"credentials-path" envconf:"GCP_CREDENTIALS_PATH"`
	}
}

var Inst = initConfig()

func initConfig() config {
	var cfg = config{}
	cfg.HomeDir = getHomeDir()
	confPostfix := os.Getenv("CONFIG_FILE_POSTFIX")
	addConfigs(&cfg, "", confPostfix)
	readAwsProfiles(&cfg)
	return cfg
}

func getHomeDir() string {
	if homeDir, err := homedir.Dir(); err == nil {
		return homeDir
	} else {
		logger.Error.Fatal(err)
		return ""
	}
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

	f, err := os.Open(getCurrentDir() + "/" + fileName)
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

func getCurrentDir() string {
	if _, filename, _, ok := runtime.Caller(0); ok {
		return path.Dir(filename)
	} else {
		logger.Error.Panic("No caller information")
		return ""
	}
}
