package controllers

import (
	"fmt"
	"github.com/BEP/bep_backend/handler"
)

type BepController struct {
	BaseController
}

func (this *BepController) CreateUser() {
	fmt.Println("call CreateUser function")
	requestData := this.JsonData()
	userId := requestData["userId"].(string)
	fmt.Println("get userId: ", userId)
	if userId == "" {
		return
	}
	msg, err := handler.App.CreateUser(userId)
	if err != nil {
		this.Error(UNKNOWN, "error when create user", err)
		return
	}
	this.Success(msg, "succeed to create user")
}

func (this *BepController) GetUserBalance() {
	fmt.Println("call GetBalance function")
	requestData := this.JsonData()
	userId := requestData["userId"].(string)
	if userId == "" {
		return
	}
	balance, err := handler.App.GetUserBalance(userId)
	if err != nil {
		this.Error(UNKNOWN, "error when get user balance", err)
		return
	}
	this.Success(balance, "succeed to get user balance")
}

func (this *BepController) PushRequest() {
	fmt.Println("call PushRequest function")
	requestData := this.JsonData()
	userId := requestData["userId"].(string)
	requirement := requestData["requirement"].(string)
	reward := requestData["reward"].(float64)
	duration := requestData["duration"].(float64)
	tranId, err := handler.App.PushRequest(userId, requirement, reward, duration)
	if err != nil {
		this.Error(UNKNOWN, "error when push request", err)
		return
	}
	this.Success(tranId, "succeed to push request")
}
func (this *BepController) PushResponse() {
	fmt.Println("call PushResponse function")
}
func (this *BepController) AcceptResponse() {
	fmt.Println("call AcceptResponse function")
}
func (this *BepController) GetAllRequests() {
	fmt.Println("call GetAllRequests function")
}
func (this *BepController) GetRequestByUserId() {
	fmt.Println("call GetRequestByUserId function")
}
func (this *BepController) GetResponseByUserId() {
	fmt.Println("call GetResponseByUserId function")
}
func (this *BepController) GetResponseByRequestId() {
	fmt.Println("call GetResponseByRequestId function")
}
