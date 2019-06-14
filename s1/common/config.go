package common

import (
	"encoding/json"
	"io/ioutil"
)

var conf *Config

type MysqlConfig struct {
	Addr     string
	User     string
	Password string
	Db       string
}

type Config struct {
	Port       int          `json:"port"`
	ServerAddr string       `json:"server_addr"`
	Mysql      *MysqlConfig `json:"mysql"`
}

func InitConf(fileName string) error {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	conf = &Config{}
	if err := json.Unmarshal(fileData, conf); err != nil {
		return err
	}

	return nil
}

func GetMysqlConf() *MysqlConfig {
	return conf.Mysql
}
