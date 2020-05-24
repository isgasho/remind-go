package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Wechat struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

type DataDb struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

type Email struct {
	User string
	Pass string
	Host string
	Port int
}
type Teng struct {
	SECRETID  string
	SecretKey string
	SDKAppID string
	TemplateID string
}

type Configuration struct {
	Wechat Wechat
	Db     DataDb
	Email  Email
	Teng   Teng
}

var config *Configuration

var once sync.Once

func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("can not open the file")
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return config
}
