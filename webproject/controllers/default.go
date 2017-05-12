package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Code"] = "11111111111111"
	c.Data["Data"] = "dsdsdsdsdsd"
	c.TplName = "index.tpl"
}
