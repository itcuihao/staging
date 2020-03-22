package db

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB       *gorm.DB
	rollback bool
}
type MysqlConfig struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

func NewDao(dbConf *MysqlConfig) *Dao {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local",
		dbConf.User, dbConf.Password, dbConf.Addr, dbConf.Db))
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(1800 * time.Second) // 腾讯云默认是3600s

	newDB := &Dao{DB: db}
	return newDB
}

func (dao *Dao) New() *Dao {
	return &Dao{DB: dao.DB.New()}
}

func InitDebug(dbConf *MysqlConfig) *Dao {

	db := NewDao(dbConf)
	db.DB.Debug()
	db.DB.SetLogger(sqlog)

	return db
}

func (dao *Dao) BeginTx() *Dao {
	return &Dao{DB: dao.DB.Begin()}
}

func (dao *Dao) CommitTx() error {
	if dao.rollback {
		return dao.DB.Rollback().Error
	}
	return dao.DB.Commit().Error
}

func (dao *Dao) Rollback() {
	dao.rollback = true
}

var sqlog *sqlLogger

type sqlLogger struct {
}

// 给sql用
func (l *sqlLogger) Print(v ...interface{}) {
	msg := gorm.LogFormatter(v...)
	if len(msg) == 5 {
		fmt.Printf("%v %v %v %v %v\n", msg[0],
			strings.Replace(fmt.Sprintf("%v", msg[1]), "\n", "", -1),
			strings.Replace(fmt.Sprintf("%v", msg[2]), "\n", "", -1),
			strings.Replace(fmt.Sprintf("%v", msg[3]), "\n", "", -1),
			strings.Replace(fmt.Sprintf("%v", msg[4]), "\n", "", -1))
	} else {
		fmt.Printf("%+v\n", msg)
	}
}
