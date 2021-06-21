package routes

import (
	"demo/controller/v1"
	_ "demo/docs"
	"demo/logger"
	"demo/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	config := cors.DefaultConfig()
	config.AllowOrigins = settings.Conf.AllowOrigins
	config.AllowCredentials = true
	r.Use(cors.New(config))
	pprof.Register(r)
	if mode == "debug" || mode == "dev" || mode == "test" {
		r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	r.POST("/v1/login", v1.Login)
	SetUserRouter(r)
	return r
}
