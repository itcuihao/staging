package common

import (
	"encoding/json"
	"io/ioutil"
)

var conf *Config

type MysqlConfig struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

type Config struct {
	Env    string       `json:"env"`
	Addr   string       `json:"addr"`
	Mysql  *MysqlConfig `json:"mysql"`
	Casbin Casbin       `json:"casbin"`
}

type Casbin struct {
	Enable bool `json:"enable"`
}

func InitConfig(cfgFile string) error {
	var err error
	conf, err = readConf(cfgFile)
	if err != nil {
		return err
	}

	return nil
}

func readConf(cfgFile string) (*Config, error) {
	f, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	data := &Config{}
	err = json.Unmarshal(f, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetEnv() string {
	return conf.Env
}

func GetAddr() string {
	return conf.Addr
}

func GetMysqlCfg() *MysqlConfig {
	return conf.Mysql
}

func GetCasbin() Casbin {
	return conf.Casbin
}
