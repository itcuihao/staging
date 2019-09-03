package ato_way

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync/atomic"
	"time"
)

var (
	atoConfig atomic.Value
	chwr      = make(chan bool, 1)
	cherr     = make(chan error, 1)
	conf      *Config
)

type Config struct {
	Filename       string
	LastModifyTime int64
	MySQL          *MySQL
}

type MySQL struct {
	Host string `json:"host"`
	DB   string `json:"db"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func NewConfig(fname string) {

	conf = &Config{
		Filename: fname,
		MySQL:    read(fname),
	}

	go conf.reload()

	go func() {
		for {
			select {
			case <-chwr:
				conf.write()
			case err := <-cherr:
				fmt.Println("err: ", err.Error())
			}
		}
	}()

	go func() {
		for {
			fmt.Printf("此时的配置：%+v\n", conf.MySQL)
			time.Sleep(time.Second * 2)
		}
	}()

	return
}

func (c *Config) reload() {
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			f, _ := os.Stat(c.Filename)
			curModifyTime := f.ModTime().Unix()
			if curModifyTime > c.LastModifyTime {
				mysql := read(c.Filename)
				if mysql != nil {
					atoConfig.Store(mysql)
					chwr <- true
				}
			}
		}
	}
}

func read(fname string) *MySQL {
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		cherr <- err
		return nil
	}

	data := new(MySQL)
	err = json.Unmarshal(f, &data)
	if err != nil {
		cherr <- err
		return nil
	}
	return data
}

func (c *Config) write() {
	data := atoConfig.Load().(*MySQL)
	if data == nil {
		return
	}
	c.MySQL = data
	c.lastTime()
}

func (c *Config) lastTime() {
	fi, _ := os.Stat(c.Filename)
	c.LastModifyTime = fi.ModTime().Unix()
}
