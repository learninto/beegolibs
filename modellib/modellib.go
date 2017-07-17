package modellib

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterMySQL(obj interface{}) {
	prefix := beego.AppConfig.String("mysqlprefix")
	orm.RegisterModelWithPrefix(prefix, obj)
}

// 获取qs
func QuerySeterMySQL(obj interface{}) (qs orm.QuerySeter) {
	o := orm.NewOrm()
	return o.QueryTable(obj)
}
