package tool

import (
	"fmt"
	"github.com/itcuihao/staging/s1/dao"
	"testing"
)

func TestNewCasbinRule(t *testing.T) {
	dao := dao.InitDebug()
	NewCasbinRule(dao, "../config/rbac_model.conf")

	//获取请求的URI
	//obj := c.Request.URL.RequestURI()
	obj := "/api/v1/user/1"
	//获取请求方法
	//act := c.Request.Method
	act := "GET"
	//获取用户的角色
	sub := "admin"

	pass, err := enforcer.Enforce(sub, obj, act)
	t.Log(err)
	if pass {
		fmt.Println("通过权限")
	} else {
		fmt.Println("权限没有通过")
	}
}
