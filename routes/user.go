package routes

import (
	"demo/controller/v1"
	"github.com/gin-gonic/gin"
)

func SetUserRouter(router *gin.Engine) {
	router.POST("/v1/logout", v1.Logout)
	userRouter := router.Group("/v1/user")
	userRouter.GET("", v1.GetUsers)
}
