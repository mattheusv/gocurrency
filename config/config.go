package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

const (
	reportConfigFile = "currency.toml"
)

var (
	//ReportConfig struct export config
	ReportConfig *report
)

type report struct {
	HTTP struct {
		Host string
		Port int
	}
	Mongo struct {
		Host     string
		User     string
		Password string
		Source   string
		Port     int
		DataBase string
	}
}

//Load load report config
func Load() {
	ReportConfig = &report{}
	if err := parse(ReportConfig); err != nil {
		log.Fatalf("erro to parse config\n %v", err)
	}

}

func parse(report *report) error {
	tomlFile, err := ioutil.ReadFile(reportConfigFile)
	if err != nil {
		return err
	}

	if _, err := toml.Decode(string(tomlFile), report); err != nil {
		return err
	}
	return nil
}
