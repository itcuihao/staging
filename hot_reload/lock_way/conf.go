package lock_way

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	Filename       string
	LastModifyTime int64
	Mt             sync.RWMutex
	MySQL          *MySQL
}

type MySQL struct {
	Host string `json:"host"`
	DB   string `json:"db"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func NewConfig(fname string) *Config {
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Println(err)
		return nil
	}
	mysql := new(MySQL)
	err = json.Unmarshal(f, mysql)
	if err != nil {
		log.Println(err)
		return nil
	}

	conf := &Config{
		Filename: fname,
		Mt:       sync.RWMutex{},
		MySQL:    mysql,
	}

	go func() {
		for {
			fmt.Printf("此时的配置：%+v\n", conf.MySQL)
			time.Sleep(time.Second * 2)
		}
	}()

	go conf.reload()

	return conf
}

func (c *Config) SetMySQL(m *MySQL) {
	c.MySQL = m
}

func (c *Config) parse() bool {
	fname, _ := os.Stat(c.Filename)
	c.LastModifyTime = fname.ModTime().Unix()

	f, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		log.Println(err)
		return false
	}

	data := new(MySQL)
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Println(err)
		return false
	}
	c.Mt.Lock()
	c.MySQL = data
	c.Mt.Unlock()
	log.Printf("data: %+v\n", c.MySQL)
	return true
}

func (c *Config) reload() {
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			f, _ := os.Stat(c.Filename)
			curModifyTime := f.ModTime().Unix()
			if curModifyTime > c.LastModifyTime {
				if c.parse() {
					log.Println("loading...")
				}
			}
		}
	}
}
