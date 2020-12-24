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
<<<<<<< HEAD
	c.Data["Email"] = "astaxie@gmail.com"*/
	c.TplName = "login.html"
=======
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/

	c.TplName = "index1.html"

	c.Data["难度"] =utils.GetDifficulty()
>>>>>>> 0866a3e7300b0c3ac29b1b198826f73dd992bdea
}
