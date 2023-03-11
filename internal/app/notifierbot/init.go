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
	Listener struct {
		Address string `yaml:"address"`
	} `yaml:"listener"`
}

var Conf Config

func init() {
	p, _ := os.Getwd()
	f, _ := os.Open(p + "/configs/notifierbot.yml")
	defer func(configFile *os.File) {
		goerr.Fatal(
			configFile.Close(),
			"[Error] notifierbot.init: Close config file",
		)
	}(f)

	d := yaml.NewDecoder(f)
	err := d.Decode(&Conf)
	goerr.Fatal(
		err,
		"[Error] notifierbot.init: Reading config",
	)
}
