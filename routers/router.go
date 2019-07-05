package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"zadmin/controllers"
)

func init() {

	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("user_id").(int)
		fmt.Println(ctx.Input.Session("user_id"))
		fmt.Println(ctx.Request.RequestURI)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)

    beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
    beego.Router("/home", &controllers.HomeController{})
    beego.Router("/users", &controllers.UserController{})
    beego.Router("/settings/?:id", &controllers.SettingController{})
    beego.Router("/settings/modal/:id", &controllers.SettingController{}, "get:GetModal")
    beego.Router("/notice", &controllers.NoticeController{})

}
