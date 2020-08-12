package main

import (
	"github.com/astaxie/beego"
	_ "test/classOne/routers"
)

func main() {
	beego.AddFuncMap("prepage", handlePrePage)
	beego.AddFuncMap("nextpage", handleNextPage)
	beego.Run()
}

func handlePrePage(index int) int {
	if index <= 1 {
		return 1
	}
	return index - 1
}

func handleNextPage(index, count int) int {
	if index == count {
		return count
	}
	return index + 1
}
