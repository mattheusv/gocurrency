package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

const (
	configFile = "currency.toml"
)

var (
	//Config struct export config
	Config *config
)

type config struct {
	HTTP struct {
		Host string
		Port int
	}
}

//Load load config
func Load() {
	Config = &config{}
	if err := parse(Config); err != nil {
		log.Fatalf("erro to parse config\n %v", err)
	}

}

func parse(c *config) error {
	tomlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	if _, err := toml.Decode(string(tomlFile), c); err != nil {
		return err
	}
	return nil
}
