package models

import (
	"github.com/astaxie/beego/orm"
)

type Admins struct{
	Id          int
	Mobile        string
	Password        string
}

type Users struct{
	Id          int
	Name        string
	Mobile        string
	Password        string
}

type Configuration struct{
	Id          int
	Key          string
	Value        string
	Desc        string
	Mark        string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(Admins),
		new(Users),
		new(Configuration),
	)
}
