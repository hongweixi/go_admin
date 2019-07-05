package controllers

import (
	"github.com/astaxie/beego/orm"
	"zadmin/models"
)

type UserController struct{
	BaseController
}

func (c *UserController) Get(){
	c.View("users/lists.html", "用户管理", true)
}

func (c *UserController) Post(){

	result := make(map[string]interface{})

	result["draw"] = c.GetString("draw", "1")
	typeValue := c.GetString("columns[0][search][value]")
	value := c.GetString("columns[1][search][value]")

	pageNum, _ := c.GetInt64("length", 10)
	start, _ := c.GetInt64("start",  0)

	o := orm.NewOrm()
	user := new(models.Users)
	qs := o.QueryTable(user)

	result["recordsTotal"], _ = qs.Count()
	if len(typeValue) > 0 && len(value) > 0 {
		switch typeValue {
			case "1": //手机号
			{
				qs = qs.Filter("mobile", value)
			}
		}
	}
	result["recordsFiltered"], _ = qs.Count();

	var users []*models.Users
	_, _ = qs.Offset(start).Limit(pageNum).All(&users)

	var Result1 []interface{}
	for _, v := range users{
		Result1 = append(Result1, []interface{}{v.Id, v.Name, v.Mobile, v.Password})
	}
	result["data"] = Result1
	c.Data["json"] = result
	c.ServeJSON()
}