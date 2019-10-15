package controllers

import (
	"fmt"
	"github.com/BEP/bep_backend/handler"
	"strconv"
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
	rewardStr := requestData["reward"].(string)
	reward, err := strconv.ParseFloat(rewardStr, 64)
	if err != nil {
		this.Error(UNKNOWN, "error when parse float64", err)
		return
	}
	expiredTime := requestData["expiredTime"].(string)
	tranId, err := handler.App.PushRequest(userId, requirement, reward, expiredTime)
	if err != nil {
		this.Error(UNKNOWN, "error when push request", err)
		return
	}
	this.Success(tranId, "succeed to push request")
}

func (this *BepController) GetAllRequests() {
	fmt.Println("call GetAllRequests function")
	requests, err := handler.App.GetAllRequests()
	if err != nil {
		this.Error(UNKNOWN, "error when get all requests", err)
		return
	}
	this.Success(requests, "succeed to get all requests")
}

func (this *BepController) GetRequestByUserId() {
	fmt.Println("call GetRequestByUserId function")
	requestData := this.JsonData()
	userId := requestData["userId"].(string)
	requests, err := handler.App.GetRequestByUserId(userId)
	if err != nil {
		this.Error(UNKNOWN, "error when get requests by userId", err)
		return
	}
	this.Success(requests, "succeed to get request by userId")
}

func (this *BepController) PushResponse() {
	fmt.Println("call PushResponse function")
	requestData := this.JsonData()
	requestId := requestData["requestId"].(string)
	userId := requestData["userId"].(string)
	answer := requestData["answer"].(string)
	tranId, err := handler.App.PushResponse(requestId, userId, answer)
	if err != nil {
		this.Error(UNKNOWN, "error when push response", err)
		return
	}
	this.Success(tranId, "succeed to push response")
}

func (this *BepController) AcceptResponse() {
	fmt.Println("call AcceptResponse function")
	requestData := this.JsonData()
	userId := requestData["userId"].(string)
	requestId := requestData["requestId"].(string)
	responseId := requestData["responseId"].(string)
	tranId, err := handler.App.AcceptResponse(userId, requestId, responseId)
	if err != nil {
		this.Error(UNKNOWN, "error when accept response", err)
		return
	}
	this.Success(tranId, "succeed to accept response")
}

func (this *BepController) GetResponseByUserId() {
	fmt.Println("call GetResponseByUserId function")
	requestData := this.JsonData()
	userId := requestData["userId"].(string)
	responses, err := handler.App.GetResponseByUserId(userId)
	if err != nil {
		this.Error(UNKNOWN, "error when get responses by userId", err)
		return
	}
	this.Success(responses, "succeed to get response by userId")
}

func (this *BepController) GetResponseByRequestId() {
	fmt.Println("call GetResponseByRequestId function")
	requestData := this.JsonData()
	requestId := requestData["requestId"].(string)
	responses, err := handler.App.GetResponseByRequestId(requestId)
	if err != nil {
		this.Error(UNKNOWN, "error when get response by requestId", err)
		return
	}
	this.Success(responses, "succeed to get response by requestId")
}

func (this *BepController) GetUserRequestAndResponses() {
	fmt.Println("call GetUserRequestAndResponses function")
	requestData := this.JsonData()
	userId := requestData["userId"].(string)
	responses, err := handler.App.GetUserRequestAndResponses(userId)
	if err != nil {
		this.Error(UNKNOWN, "error when get user requests and responses", err)
		return
	}
	this.Success(responses, "succeed to get user requests and responses")
}
