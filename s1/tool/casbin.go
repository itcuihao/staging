package tool

import (
	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/dao"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

var (
	enforcer *casbin.Enforcer
)

func GetCasbinRule() *casbin.Enforcer {
	return enforcer
}

func NewCasbinRule(dao *dao.Dao, path string) {
	//path "config/rbac_model.conf"
	a, err := gormadapter.NewAdapterByDB(dao.NewDB())
	if err != nil {
		common.Log.Errorf("casbin error: %v", err)
		panic(err)
	}
	enforcer, err = casbin.NewEnforcer(path, a)
	if err != nil {
		common.Log.Errorf("casbin error: %v", err)
		panic(err)
	}
	if err := enforcer.LoadPolicy(); err != nil {
		common.Log.Errorf("casbin error: %v", err)
		panic(err)
	}
}
