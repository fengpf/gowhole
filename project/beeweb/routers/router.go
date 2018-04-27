package routers

import (
	"github.com/astaxie/beego"
	"gostudy/webproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
