package mysql

import (
	"demo/settings"
	"fmt"
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
	if err := Init(&cfg); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
}

func TestGetUser(t *testing.T) {

}
