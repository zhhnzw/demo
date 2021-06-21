package mysql

import (
	"demo/settings"
	"fmt"
	"testing"
)

// 数据库层的单元测试需要初始化
func init() {
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	if err := Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
}

func TestGetUser(t *testing.T) {

}
