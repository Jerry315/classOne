package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/classOne/models"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) ShowReg() {
	u.TplName = "register.html"
}

func (u *UserController) HandleReg() {
	// 获取表单数据
	name := u.GetString("userName")
	pwd := u.GetString("password")
	// 验证数据不能为空
	if name == "" && pwd == "" {
		fmt.Println("用户名或密码不能为空")
		u.TplName = "register.html"
		return
	}

	// 获取orm对象
	o := orm.NewOrm()
	//初始化对象
	user := models.User{Name: name, PassWord: pwd}
	// 插入数据
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Printf("插入数据失败：%v\n", err)
	}
	u.Redirect("/login", 302)
}

func (u *UserController) ShowLogin() {
	u.TplName = "login.html"
}

func (u *UserController) HandleLogin() {
	// 获取表单数据
	name := u.GetString("userName")
	pwd := u.GetString("password")
	// 验证数据不能为空
	if name == "" && pwd == "" {
		fmt.Println("用户名或密码不能为空")
		u.TplName = "login.html"
		return
	}

	// 获取orm对象
	o := orm.NewOrm()
	//初始化对象
	user := models.User{Name: name}
	err := o.Read(&user, "Name")
	if err != nil {
		fmt.Println("用户不存在")
		u.Data["errmsg"] = "用户不存在"
		u.TplName = "login.html"
		return
	}

	if user.PassWord != pwd {
		fmt.Println("密码错误")
		u.Data["errmsg"] = "密码错误"
		u.TplName = "login.html"
		return
	}
	u.Redirect("/showArticleList", 302)

}
