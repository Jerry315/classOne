package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"test/classOne/models"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) ShowArticleList() {
	// 获取数据，高级查询，指定表
	o := orm.NewOrm()
	qs := o.QueryTable("Article") // query string
	articles := []models.Article{}

	// 查询文章总数
	count, _ := qs.RelatedSel("ArticleType").Count()
	// 每页显示条目
	pageSize := 2
	// 根据每页显示条目，计算得出总的页码
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	// 获取前端传递页码，如果没传，默认页码为1
	pageIndex, err := this.GetInt("pageIndex")
	if err != nil {
		pageIndex = 1
	}

	// 获取分页数据
	start := (pageIndex - 1) * pageSize
	// 数据库中按照分页获取数据
	qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articles)

	var articleTypes []*models.ArticleType
	o.QueryTable("article_type").All(&articleTypes)

	this.Data["pageIndex"] = pageIndex
	this.Data["pageCount"] = int(pageCount)
	this.Data["count"] = count
	this.Data["articles"] = articles
	this.Data["articleTypes"] = articleTypes
	this.TplName = "index.html"
}

//展示添加文章页面
func (this *ArticleController) ShowAddArticle() {
	var articleTypes []*models.ArticleType
	o := orm.NewOrm()
	o.QueryTable("article_type").All(&articleTypes)
	this.Data["articleTypes"] = articleTypes
	this.TplName = "add.html"
}

func (this *ArticleController) HandleAddArticle() {
	// 获取数据
	title := this.GetString("articleName")
	content := this.GetString("content")
	articleTypeId,_ := this.GetInt("select")
	// 校验数据
	if title == "" && content == "" {
		this.Data["errmsg"] = "标题或内容不能为空"
		this.TplName = "add.html"
		return
	}
	// 处理文件上传
	file, head, err := this.GetFile("uploadname")
	defer file.Close()
	if err != nil {
		this.Data["errmsg"] = "文件上传失败"
		this.TplName = "add.html"
		return
	}
	// 文件大小限制，5000000
	if head.Size > 5000000 {
		this.Data["errmsg"] = "文件太大，请重新上传"
		this.TplName = "add.html"
	}
	// 文件格式限制.jpg .png .jpeg
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		this.Data["errmsg"] = "文件格式错误。请重新上传"
		this.TplName = "add.thml"
		return
	}
	// 防止重名
	fileName := time.Now().Format("2006-01-02-15-16-14") + ext
	fmt.Println("上传文件：", fileName)
	// 处理数据，插入操作
	err = this.SaveToFile("uploadname", "./static/img/"+fileName)
	if err != nil {
		fmt.Printf("文件保存失败，err: %v\n", err)

	}
	// 返回页面
	o := orm.NewOrm()

	var articleType models.ArticleType
	articleType.Id = articleTypeId
	o.Read(&articleType)

	var article models.Article
	article.ArticleName = title
	article.AccountTent = content
	article.ArticleType = &articleType
	article.Img = "../static/img/" + fileName
	_, err = o.Insert(&article)
	if err != nil {
		fmt.Printf("新增文章失败，%v\n", err)
	}
	this.Redirect("/showArticleList", 302)
	//this.Ctx.WriteString("添加文章成功")
}

func (this *ArticleController) ShowArticleDetail() {
	// 获取文章id
	aid, err := this.GetInt("articleId")
	if err != nil {
		this.Ctx.WriteString("文章id不存在")
		return
	}

	// 根据文章id从数据库中获取文章
	o := orm.NewOrm()
	var article models.Article
	var articleType models.ArticleType
	article.Id = aid
	err = o.Read(&article, "Id")
	if err != nil {
		this.Ctx.WriteString("数据库中查无此数据")
		return
	}

	articleType.Id = article.ArticleType.Id
	o.Read(&articleType)
	Tname := articleType.Tname

	// 修改阅读量
	article.Account += 1
	o.Update(&article)

	// 渲染视图页面
	this.Data["article"] = article
	this.Data["Tname"] = Tname
	this.TplName = "content.html"
}

func (this *ArticleController) DelArticle() {
	// 获取文章id
	aid, err := this.GetInt("articleId")
	if err != nil {
		this.Ctx.WriteString("文章id不存在")
		return
	}

	// 根据问文章id从数据库中获取文章
	o := orm.NewOrm()
	var article models.Article
	article.Id = aid
	o.Delete(&article)
	this.Redirect("/showArticleList", 302)
}

func (this *ArticleController) ShowAddArticleType() {
	var articleTypes []*models.ArticleType
	o := orm.NewOrm()
	o.QueryTable("article_type").All(&articleTypes)
	this.Data["articleTypes"] = articleTypes
	this.TplName = "addType.html"
}

func (this *ArticleController) AddArticleType() {
	typeName := this.GetString("typeName")
	var typeArticle models.ArticleType
	typeArticle.Tname = typeName

	o := orm.NewOrm()
	_, err := o.Insert(&typeArticle)
	if err != nil {
		fmt.Printf("插入文章类型失败，%v\n", err)
	}
	this.Redirect("/addArticleType", 302)
}
