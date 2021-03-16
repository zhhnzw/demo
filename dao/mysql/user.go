package mysql

type User struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}
