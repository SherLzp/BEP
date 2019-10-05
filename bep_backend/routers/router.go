package routers

import (
	"github.com/BEP/bep_backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/createUser", &controllers.BepController{}, "post:CreateUser")
	beego.Router("/getUserBalance", &controllers.BepController{}, "post:GetUserBalance")
	beego.Router("/pushRequest", &controllers.BepController{}, "post:PushRequest")
	beego.Router("/pushResponse", &controllers.BepController{}, "post:PushResponse")
	beego.Router("/acceptResponse", &controllers.BepController{}, "post:AcceptResponse")
	beego.Router("/getAllRequests", &controllers.BepController{}, "get:GetAllRequests")
	beego.Router("/getRequestByUserId", &controllers.BepController{}, "post:GetRequestByUserId")
	beego.Router("/getResponseByUserId", &controllers.BepController{}, "post:GetResponseByUserId")
	beego.Router("/getResponseByRequestId", &controllers.BepController{}, "post:GetResponseByRequestId")
}
