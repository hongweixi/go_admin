package controllers

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"zadmin/models"
)

type SettingController struct{
	BaseController
}

func (c *SettingController) Get(){
	var Configurations []*models.Configuration
	o := orm.NewOrm()
	Configuration := new(models.Configuration)
	_, _ = o.QueryTable(Configuration).All(&Configurations)

	c.Data["lists"] = Configurations
	c.View("system/settings.html", "系统设置", true)
}

func (c *SettingController) GetModal(){

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	o := orm.NewOrm()
	Configuration := models.Configuration{Id: idInt}
	_ = o.Read(&Configuration)

	c.Data["Configuration"] = Configuration
	c.TplName = "system/settings_modal.html"
}

func (c *SettingController) Post(){
	mark := c.GetString("mark", "")
	key := c.GetString("key")
	value := c.GetString("value")
	desc := c.GetString("desc", "")
	if key == "" || value == "" {
		c.Data["json"] = map[string]interface{}{"status":2, "message":"请输入配置值"}
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	Configuration := new(models.Configuration)
	Configuration.Key = key
	Configuration.Value = value
	Configuration.Mark = mark
	Configuration.Desc = desc
	_, err := o.Insert(Configuration)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"status":2, "message":"新增失败"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"status":1, "message":"新增成功"}
	c.ServeJSON()
	return
}

func (c *SettingController) Delete(){
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	if idInt == 0{
		c.Data["json"] = map[string]interface{}{"status":2, "message":"参数错误"}
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	if _, err := o.Delete(&models.Configuration{Id: idInt}); err == nil {
		c.Data["json"] = map[string]interface{}{"status":1, "message":"删除成功"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"status":2, "message":"删除失败"}
	c.ServeJSON()
	return
}

func (c *SettingController) Put(){
	mark := c.GetString("mark", "")
	key := c.GetString("key")
	value := c.GetString("value")
	desc := c.GetString("desc", "")
	if key == "" || value == "" {
		c.Data["json"] = map[string]interface{}{"status":2, "message":"请输入配置值"}
		c.ServeJSON()
		return
	}
	id, _ := c.GetInt("id")
	if id == 0{
		c.Data["json"] = map[string]interface{}{"status":2, "message":"参数错误"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	Configuration := models.Configuration{Id: id}
	if o.Read(&Configuration) == nil {
		Configuration.Key = key
		Configuration.Value = value
		Configuration.Mark = mark
		Configuration.Desc = desc


		if _, err := o.Update(&Configuration); err == nil {
			c.Data["json"] = map[string]interface{}{"status":1, "message":"更新成功"}
			c.ServeJSON()
			return
		}
	}
	c.Data["json"] = map[string]interface{}{"status":2, "message":"更新失败"}
	c.ServeJSON()
	return
}