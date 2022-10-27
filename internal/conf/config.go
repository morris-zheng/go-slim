package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Env   string `yaml:"env"`
	Debug bool   `yaml:"debug"`
	Port  int    `yaml:"port"`
	Mysql Mysql  `yaml:"mysql"`
}

func Load(path string) *Config {
	cb, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("read config err", err)
	}
	var c Config
	err = yaml.Unmarshal(cb, &c)
	if err != nil {
		log.Fatal("unmarshal config err", err)
	}
	return &c
}
