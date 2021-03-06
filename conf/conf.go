package conf

import (
	toml "github.com/pelletier/go-toml"
)

type Config struct {
	JSONFileDir    string
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

	return &conf, nil
}
