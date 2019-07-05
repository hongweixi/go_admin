package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct{
	beego.Controller
}

func (c *BaseController) View(html string, title string, haveJs bool){
	c.Layout = "layout/base.html"
	c.TplName = html
	c.Data["Title"] = title
	c.Data["Mobile"] = c.GetSession("user_mobile")
	c.Data["Menus"] = MenuGet()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "layout/head.html"
	c.LayoutSections["ScriptsBase"] = "layout/js_base.html"
	c.LayoutSections["Scripts"] = ""
	if haveJs{
		c.LayoutSections["Scripts"] = c.parseScriptsPath(html)
	}
}

func (c *BaseController) parseScriptsPath(tplName string) string{
	viewSplit := strings.Split(tplName, "/")
	if len(viewSplit) == 2{
		jsSplit := strings.Split(viewSplit[1], ".")
		return viewSplit[0] + "/" + jsSplit[0] + "_js." + jsSplit[1]
	}else {
		jsSplit := strings.Split(tplName, ".")
		return jsSplit[0] + "_js." + jsSplit[1]
	}
}