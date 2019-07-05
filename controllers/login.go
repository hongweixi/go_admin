package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zadmin/models"
)

type LoginController struct{
	beego.Controller
}

func (c *LoginController) Get(){
	c.TplName = "login/login.html"
	c.Data["Title"] = "后台登录"
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}


func (c *LoginController) Post(){
	mobile := c.GetString("mobile")
	password := c.GetString("password")

	o := orm.NewOrm()
	admin := models.Admins{Mobile: mobile}
	err := o.Read(&admin, "Mobile")

	if err == orm.ErrNoRows {
		c.Data["json"] = map[string]interface{}{"status":2, "message":"登录失败，账号不存在"}
		c.ServeJSON()
	}
	if admin.Password != password{
		c.Data["json"] = map[string]interface{}{"status":2, "message":"登录失败，密码错误"}
		c.ServeJSON()
	}
	//存储session
	c.SetSession("user_id", admin.Id)
	c.SetSession("user_mobile", admin.Mobile)

	c.Data["json"] = map[string]interface{}{"status":1, "message":"登录成功"}
	c.ServeJSON()
}