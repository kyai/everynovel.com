package controllers

import (
	"github.com/astaxie/beego"
)

type WordController struct {
	beego.Controller
}

func (c *WordController) Get() {
	c.TplName = "index.tpl"
}
