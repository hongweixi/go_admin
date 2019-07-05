package controllers

/*type MenuController struct{
	BaseController
}*/

type MenuStruct struct {
	Header string
	Key string
	Icon string
	Menu []ChildMenuStruct
}

type ChildMenuStruct struct {
	Link string
	Name string
	Key string
	Submenus []Child2MenuStruct
}

type Child2MenuStruct struct {
	Link string
	Name string
	Key string
	Submenus []struct{}
}

func MenuGet() []MenuStruct {
	menus := []MenuStruct{
		{
			"用户管理",
			"user_section",
			"fa fa-user-circle-o",
			[]ChildMenuStruct{
				{"/users", "用户管理", "user_list", []Child2MenuStruct{}},
			},
		},
		{
			"系统设置",
			"system_section",
			"fa fa-cog",
			[]ChildMenuStruct{
				{"/settings", "系统设置", "settings", []Child2MenuStruct{}},
				{"/notice", "公告中心", "notice", []Child2MenuStruct{}},
			},
		},
		{
			"团队管理",
			"team_section",
			"fa fa-users",
			[]ChildMenuStruct{
				{"/users_structure", "结构图", "users_structure", []Child2MenuStruct{}},
			},
		},

	}

	return menus
}