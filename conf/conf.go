package conf

import (
	"time"

	toml "github.com/pelletier/go-toml"
)

type Config struct {
	JSONFileDir    string
	ScanPeriod     time.Duration
	ReceiverLat    float64
	ReceiverLon    float64
	ReceiverHeight float64
}

func LoadConfig() (*Config, error) {
	var confTree *toml.Tree
	var err error
	if confTree, err = toml.LoadFile("config.toml"); err != nil {
		return nil, err
	}

	conf := Config{}
	if err := confTree.Unmarshal(&conf); err != nil {
		return nil, err
	}

	conf.ScanPeriod = conf.ScanPeriod * time.Second
	return &conf, nil
}
