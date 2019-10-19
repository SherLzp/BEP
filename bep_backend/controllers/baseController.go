package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

const (
	SUCCESS       = 0
	UNKNOWN       = 1
	PARSEERROR    = 98
	SESSIONEXPIRE = 99
)

// General Response Data
type ResResult struct {
	Status  int         `json:"status"` // response status 200/400/500...
	Data    interface{} `json:"data"`   // response data
	Message string      `json:"msg"`    // response message
}

// provide basic operations
type BaseController struct {
	beego.Controller
}

// parse requestBody's data
//func (this *BaseController) ParseBody(v interface{}) error {
//	return json.Unmarshal(this.Ctx.Input.RequestBody, v)
//}

func (this *BaseController) RequestBody() []byte {
	return this.Ctx.Input.RequestBody
}

func (this *BaseController) decodeRawRequestBodyJson() map[string]interface{} {
	var mm interface{}
	requestBody := make(map[string]interface{})
	json.Unmarshal(this.RequestBody(), &mm)
	if mm != nil {
		var m1 map[string]interface{}
		m1 = mm.(map[string]interface{})
		for k, v := range m1 {
			requestBody[k] = v
		}
	}
	return requestBody
}

func (this *BaseController) JsonData() map[string]interface{} {
	return this.decodeRawRequestBodyJson()
}

// response parse request data error
func (this *BaseController) ResParseError(err error) {
	this.Error(PARSEERROR, "the format of request data is wrong！", err)
}

// response json data
func (this *BaseController) ResJson(v interface{}) {
	this.Data["json"] = v
	this.ServeJSON()
}

// response when succeed
func (this *BaseController) Success(data interface{}, msg string) {
	result := ResResult{
		Status:  SUCCESS,
		Data:    data,
		Message: msg,
	}
	this.ResJson(result)
}

// response when error
func (this *BaseController) Error(status int, msg string, err error) {
	result := ResResult{
		Status:  status,
		Message: msg,
	}
	fmt.Println(fmt.Sprintf("%s:%v", msg, err))
	//beego.Error(fmt.Sprintf("%s:%v", msg, err))
	this.ResJson(result)
}
