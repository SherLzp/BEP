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
