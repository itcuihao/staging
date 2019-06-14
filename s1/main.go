package main

import (
	"flag"
	"net/http"

	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/handle"
	"github.com/itcuihao/staging/s1/storage"
)

func main() {
	userHandle := handle.NewUserHandle(db)

	http.HandleFunc("/user/:id", userHandle.GetUser)
	http.ListenAndServe(":8080", nil)
}

var (
	db *storage.DB
)

func init() {
	var confFile string
	flag.StringVar(&confFile, "c", "config/dev.json", "conf file")
	flag.Parse()

	if err := common.InitConf(confFile); err != nil {
		panic(err)
	}

	db = storage.NewDB(common.GetMysqlConf())
}
