package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

func init() {
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
