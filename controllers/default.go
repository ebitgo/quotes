package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "ebitgo.com"
	c.Data["Email"] = "support@ebitgo.com"
	c.TplName = "index.tpl"
}
