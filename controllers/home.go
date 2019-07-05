package controllers

type HomeController struct{
	BaseController
}

func (c *HomeController) Get(){
	c.View("index.html", "仪表盘", false)
}


