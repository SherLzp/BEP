package routers

import (
	"github.com/BEP/bep_backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/createUser", &controllers.BepController{}, "post:CreateUser")
	beego.Router("/getUserBalance", &controllers.BepController{}, "post:GetUserBalance")
}
