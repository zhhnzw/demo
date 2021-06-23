package v1

import (
	"demo/dao/mysql"
	"demo/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 需要用数据库的单元测试需要对其初始化
func init() {
	cfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "root",
		DbName:       "test",
		Port:         3306,
		MaxOpenConns: 200,
		MaxIdleConns: 50,
	}
	if err := mysql.Init(&cfg); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default() // 如果直接引用 routers 模块，会造成循环引用，所以在此初始化一个 default
	url := "/v1/user"
	r.GET(url, GetUsers)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
