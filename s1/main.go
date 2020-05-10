package main

import (
	"context"
	"flag"
	"github.com/itcuihao/staging/s1/tool"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/itcuihao/staging/s1/common"
	"github.com/itcuihao/staging/s1/dao"
	"github.com/itcuihao/staging/s1/handle"
	"github.com/itcuihao/staging/s1/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	confFile string
	db       *dao.Dao

	userHandle *handle.UserHandle
	roleHandle *handle.RoleHandle
)

func main() {
	flag.Parse()
	if err := common.InitConfig(confFile); err != nil {
		panic(err)
	}
	db = dao.NewDao(common.GetMysqlCfg())

	tool.NewCasbinRule(db, "config/rbac_model.conf")

	r := gin.New()

	// 不记录日志，不添加认证
	r.HEAD("/api/health", HealthHandler)
	r.GET("/api/health", HealthHandler)
	pprof.Register(r, "api/pprof")

	r.Use(middlewares.AccessMiddleware(), gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{http.MethodGet, http.MethodOptions, http.MethodPost, http.MethodDelete},
		AllowHeaders:    []string{"*"},
	}))

	newHandle(db)
	newRouterV1(r)

	srv := &http.Server{
		Addr:    common.GetAddr(),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			common.Log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	common.Log.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		common.Log.Fatalf("Server Shutdown:", err)
	}
	common.Log.Infof("Server exiting")
}

func init() {
	flag.StringVar(&confFile, "c", "config/dev.json", "conf file")
}

func newHandle(db *dao.Dao) {
	userHandle = handle.NewUserHandle(db)
	roleHandle = handle.NewRoleHandle(db)
}

func newRouterV1(r *gin.Engine) {
	mid := middlewares.Middleware{
		Skippers: []middlewares.SkipperFunc{
			middlewares.AllowPathPrefixSkipper("api/v1/login"),
		},
		CasbinRule: tool.GetCasbinRule(),
	}
	r.Use(mid.AuthMiddleware, mid.CasbinMiddleware)
	v1 := r.Group("/api/v1")
	{
		login := v1.Group("login")
		{
			login.POST("", userHandle.Login)
		}
		user := v1.Group("user")
		{
			user.GET(":id", userHandle.GetUser)
		}
		role := v1.Group("role")
		{
			role.GET("", roleHandle.GetRole)
		}
	}
}

func HealthHandler(c *gin.Context) {
	// 健康检查
	c.String(http.StatusOK, "ok")
}
