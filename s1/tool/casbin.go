package tool

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/dao"
)

var (
	enforcer *casbin.Enforcer
)

func GetCasbinRule() *casbin.Enforcer {
	return enforcer
}

func NewCasbinRule(dao *dao.Dao) {
	a, err := gormadapter.NewAdapterByDB(dao.NewDB())
	if err != nil {
		common.Log.Errorf("casbin error: %v", err)
		return
	}
	enforcer, err = casbin.NewEnforcer("../config/rbac_model.conf", a)
	if err != nil {
		common.Log.Errorf("casbin error: %v", err)
		return
	}
	enforcer.LoadPolicy()
}
