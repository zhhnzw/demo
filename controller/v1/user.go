package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
	Code    string      `json:"code"`
}

// Login 登录接口
// @Summary 登录接口
// @Description 登录接口
// @Tags 登录
// @Accept application/json
// @Produce application/json
// @Param object query mysql.User false "查询参数"
// @Success 200 {object} Resp
// @Router /v1/login [post]
func Login(c *gin.Context) {
	resp := Resp{Data: make(map[string]string), Code: "1"}
	c.JSON(http.StatusOK, resp)
}

// Logout 注销接口
// @Summary 注销接口
// @Description 注销接口
// @Tags 注销
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "用户令牌"
// @Param object query mysql.User false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} Resp
// @Router /v1/logout [post]
func Logout(c *gin.Context) {
	resp := Resp{Data: make(map[string]string), Code: "1"}
	c.JSON(http.StatusOK, resp)
}

// GetUsers 获取用户信息接口
// @Summary 获取用户信息接口
// @Description 获取用户信息接口
// @Tags 获取用户信息
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "用户令牌"
// @Param object query mysql.User false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} Resp
// @Router /v1/user [get]
func GetUsers(c *gin.Context) {
	resp := Resp{Data: make(map[string]string), Code: "1"}
	c.JSON(http.StatusOK, resp)
}
