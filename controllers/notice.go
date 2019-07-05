package controllers

type NoticeController struct{
	BaseController
}

func (c *NoticeController) Get(){

	c.View("users/notice.html", "公告管理", true)
}

func (c *NoticeController) Post(){


}