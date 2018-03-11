package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

type urlConfig struct {
	Domian      string
	Signin      string
	Signup      string
	CreateUser  string
	DeleteImage string
	Logout      string
	ViewImage   string
}

var URL urlConfig

func init() {
	URL.Domian = "http://127.0.0.1:7070"
	URL.Signin = "/v1/user/login"
	URL.Signup = "/v1/user/signup"
	URL.CreateUser = "/v1/user"
	URL.DeleteImage = "/v1/image"
	URL.Logout = "/v1/user/logout"
	URL.ViewImage = "/v1/image/view"

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(User), new(Image))
	orm.RegisterDataBase("default", "mysql", "root:123456@/pic_bed?charset=utf8")
	if err := orm.RunSyncdb("default", true, true); err != nil {
		panic(err)
	}

	O = orm.NewOrm()
	O.Using("default")
}
