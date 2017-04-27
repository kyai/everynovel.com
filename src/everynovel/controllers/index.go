package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	uid := c.GetSession("uid")
	if uid != nil{
		c.Data["uid"] = c.GetSession("uid")
		c.Data["uname"] = c.GetSession("uname")
	}

	c.TplName = "index.html"
}
