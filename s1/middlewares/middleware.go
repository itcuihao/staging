package middlewares

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Skippers   []SkipperFunc
	CasbinRule *casbin.Enforcer
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)
		fmt.Println(path)
		fmt.Println(pathLen)
		if pathLen > 0 {
			// 使/api/v1/为api/v1
			path = path[1:]
			pathLen = len(path)
			fmt.Println(path)
			fmt.Println(pathLen)
		}
		for _, p := range prefixes {
			fmt.Println(p, " ", path)
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// SkipHandler 统一处理跳过函数
func SkipHandler(c *gin.Context, skippers ...SkipperFunc) bool {
	for _, skipper := range skippers {
		fmt.Println(skipper)
		if skipper(c) {
			return true
		}
	}
	fmt.Println(len(skippers))
	return false
}

func EmptyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
