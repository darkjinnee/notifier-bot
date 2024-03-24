package notifierbot

import (
	goerr "github.com/darkjinnee/go-err"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`
	Api struct {
		URL   string `yaml:"url"`
		Token string `yaml:"token"`
	} `yaml:"api"`
	Http struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Bot struct {
		Token   string `yaml:"token"`
		Timeout int    `yaml:"timeout"`
		Debug   bool   `yaml:"debug"`
	} `yaml:"bot"`
}

var Conf Config

func init() {
	p, _ := os.Getwd()
	f, _ := os.Open(p + "/configs/notifierbot.yml")
	defer func(configFile *os.File) {
		goerr.Fatal(
			configFile.Close(),
			"[Error] notifierbot.init: Failed close config file",
		)
	}(f)

	d := yaml.NewDecoder(f)
	err := d.Decode(&Conf)
	goerr.Fatal(
		err,
		"[Error] notifierbot.init: Failed reading config",
	)
}
