package notifierbot

import (
	"github.com/darkjinnee/go-err"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`
}

var conf Config

func init() {
	dirPath, _ := os.Getwd()
	configFile, _ := os.Open(dirPath + "/configs/notifierbot.yml")
	defer func(configFile *os.File) {
		goerr.Fatal(
			configFile.Close(),
			"[Error] notifierbot.init: Close config file",
		)
	}(configFile)

	decoder := yaml.NewDecoder(configFile)
	err := decoder.Decode(&conf)
	goerr.Fatal(
		err,
		"[Error] notifierbot.init: Reading config",
	)
}
