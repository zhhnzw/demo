package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
