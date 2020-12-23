package controllers

import (
	"github.com/astaxie/beego"
	"manypeoplework/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/

	c.TplName = "index1.html"

	c.Data["难度"] =utils.GetDifficulty()
}
