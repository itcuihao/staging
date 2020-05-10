package tool

import (
	"github.com/itcuihao/staging/s1/dao"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	dao := dao.InitDebug()
	AutoMigrate(dao)
}
