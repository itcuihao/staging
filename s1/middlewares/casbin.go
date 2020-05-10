package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itcuihao/staging/s1/common"
	"net/http"
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
	if b, err := m.CasbinRule.Enforce(sub, obj, act); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  common.ErrCodePermissionDenied,
			"error": "没有权限访问",
		})
		c.Abort()
		return
	} else if !b {
		fmt.Println(b)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  common.ErrCodePermissionDenied,
			"error": "没有权限访问",
		})
		c.Abort()
		return
	}
	c.Next()
}
