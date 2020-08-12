package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int        `orm:"pk;auto"`
	Name     string     `orm:"unique"`
	PassWord string     `orm:"size(20)"`
	Articles []*Article `orm:"rel(m2m)"` // 设置多对多关系
}

type Article struct {
	Id          int       `orm:"pk;auto"`
	ArticleName string    `orm:"size(20)"`
	Atime       time.Time `orm:"auto_now"`
	Account     int       `orm:"default(0);null"`
	AccountTent string    `orm:"size(500)"`
	Img         string    `orm:"size(100)"`

	ArticleType *ArticleType `orm:"rel(fk)"`       // 设置一对多关系
	Users       []*User      `orm:"reverse(many)"` // 设置多对多的反向关系
}

// 类型表
type ArticleType struct {
	Id       int
	Tname    string     `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"` // 设置一对多的反向关系
}

func init() {
	// 获取连接对象
	_ = orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/web?charset=utf8")
	// 创建表
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	// 生成表
	_ = orm.RunSyncdb("default", false, true)
}
