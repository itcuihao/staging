package dao

import (
	"flag"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/itcuihao/staging/s1/common"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	db       *gorm.DB
	rollback bool
}

func NewDao(dbConf *common.MysqlConfig) *Dao {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local",
		dbConf.User, dbConf.Password, dbConf.Addr, dbConf.Db))
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(1800 * time.Second) // 腾讯云默认是3600s

	db.LogMode(true)
	newDB := &Dao{db: db}
	return newDB
}

func (db *Dao) New() *Dao {
	return &Dao{db: db.db.New()}
}

func InitDebug() *Dao {
	var debugConfFile string
	flag.StringVar(&debugConfFile, "c", "../config/dev.json", "conf file")
	flag.Parse()

	if err := common.InitConfig(debugConfFile); err != nil {
		panic(err)
	}
	db := NewDao(common.GetMysqlCfg())

	//db.db.SetLogger(sqlog)

	return db
}
func InitExplain() *Dao {
	var debugConfFile string
	flag.StringVar(&debugConfFile, "c", "../config/config.json", "conf file")
	flag.Parse()

	if err := common.InitConfig(debugConfFile); err != nil {
		panic(err)
	}
	db := NewDao(common.GetMysqlCfg())
	db.db.Debug()

	//db.db.SetLogger(sqlog)

	return db
}

func (db *Dao) NewDB() *gorm.DB {
	return db.db.New()
}

func (db *Dao) BeginTx() *Dao {
	return &Dao{db: db.db.Begin()}
}

func (db *Dao) CommitTx() error {
	if db.rollback {
		return db.db.Rollback().Error
	}
	return db.db.Commit().Error
}

func (db *Dao) Rollback() {
	db.rollback = true
}

var sqlog *sqlLogger

type sqlLogger struct {
}

// 给sql用
func (l *sqlLogger) Print(v ...interface{}) {
	msg := gorm.LogFormatter(v...)
	if len(msg) == 5 {
		common.Log.Infof("%v %v %v %v %v", msg[0],
			strings.Replace(fmt.Sprintf("%v", msg[1]), "\n", "", -1),
			strings.Replace(fmt.Sprintf("%v", msg[2]), "\n", "", -1),
			strings.Replace(fmt.Sprintf("%v", msg[3]), "\n", "", -1),
			strings.Replace(fmt.Sprintf("%v", msg[4]), "\n", "", -1))
	} else {
		common.Log.Infof("%+v", msg)
	}
}
