package routers

import (
	"github.com/astaxie/beego"
	"test/classOne/controllers"
)

func init() {
	beego.Router("/register", &controllers.UserController{}, "get:ShowReg;post:HandleReg")
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/showArticleList", &controllers.ArticleController{}, "get:ShowArticleList")
	beego.Router("/addArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/showArticleDetail", &controllers.ArticleController{}, "get:ShowArticleDetail")
	beego.Router("/addArticleType",&controllers.ArticleController{},"get:ShowAddArticleType;post:AddArticleType")
	beego.Router("/delArticle",&controllers.ArticleController{},"post:DelArticle")
}
