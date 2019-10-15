package routers

import (
	"github.com/BEP/bep_backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/createUser", &controllers.BepController{}, "post:CreateUser")
	beego.Router("/getUserBalance", &controllers.BepController{}, "post:GetUserBalance")
	beego.Router("/request/pushRequest", &controllers.BepController{}, "post:PushRequest")
	beego.Router("/response/pushResponse", &controllers.BepController{}, "post:PushResponse")
	beego.Router("/user/acceptResponse", &controllers.BepController{}, "post:AcceptResponse")
	beego.Router("/request/getAllRequests", &controllers.BepController{}, "get:GetAllRequests")
	beego.Router("/request/getRequestByUserId", &controllers.BepController{}, "post:GetRequestByUserId")
	beego.Router("/response/getResponseByUserId", &controllers.BepController{}, "post:GetResponseByUserId")
	beego.Router("/response/getResponseByRequestId", &controllers.BepController{}, "post:GetResponseByRequestId")
	beego.Router("/response/getUserRequestAndResponses", &controllers.BepController{}, "post:GetUserRequestAndResponses")
}
