package middlewares

import (
	"fmt"
	"net/http"

	"github.com/itcuihao/staging/s1/common"

	"github.com/gin-gonic/gin"
)

// CasbinMiddleware casbin中间件
func (m Middleware) CasbinMiddleware(c *gin.Context) {
	cfg := common.GetCasbin()
	if !cfg.Enable {
		c.Next()
		return
	}
	if SkipHandler(c, m.Skippers...) {
		c.Next()
		return
	}
	// 资源
	obj := c.Request.URL.Path
	// 方法
	act := c.Request.Method
	// 用户
	sub := c.GetString(AuthKeyUserRole)
	fmt.Println(sub, act, obj)
	pass, err := m.CasbinRule.Enforce(sub, obj, act)
	fmt.Println(pass)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  common.ErrCodePermissionDenied,
			"error": "没有权限访问",
		})
		c.Abort()
		return
	} else if !pass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  common.ErrCodePermissionDenied,
			"error": "没有权限访问",
		})
		c.Abort()
		return
	}
	c.Next()
}
