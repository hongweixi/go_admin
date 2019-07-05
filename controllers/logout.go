package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct{
	beego.Controller
}

func (c *LogoutController) Get(){
	userId := c.GetSession("user_id")
	if userId != nil {
		c.DelSession("user_id")
		c.DelSession("user_mobile")
	}
	c.Redirect("/login", 302)
}