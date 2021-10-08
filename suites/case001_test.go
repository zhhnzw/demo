package suites

import (
	v1 "demo/controller/v1"
	"demo/dao/mysql"
	"demo/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
)

type Case001TestSuite struct {
	suite.Suite
	router *gin.Engine
}

// 需要用数据库的单元测试需要对其初始化
func (suite *Case001TestSuite) SetupSuite() {
	fmt.Println("SetupSuite begin")
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
		suite.FailNow("init mysql failed:%v", err)
		return
	}
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default() // 如果直接引用 routers 模块，会造成循环引用，所以在此初始化一个 default
	fmt.Println("SetupSuite finished")
}

func (suite *Case001TestSuite) TearDownSuite() {
	mysql.Close()
}

func (suite *Case001TestSuite) TestGetUsers() {
	url := "/v1/user"
	fmt.Printf("%v\n", suite.router)
	suite.router.GET(url, v1.GetUsers)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	suite.Nil(err)
	suite.Equal(200, w.Code)
}
