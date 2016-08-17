package routers

import (
	"github.com/astaxie/beego"
	"wwc_register/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/adduser", &controllers.RegController{}, "post:AddUser")
}
