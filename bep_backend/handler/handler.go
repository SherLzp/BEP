package handler

import (
	"encoding/json"
	"github.com/BEP/bep_backend/service"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

type Application struct {
	beego.Controller
	Setup *service.ServiceSetup
}

var App Application

func (app *Application) CreateUser(userId string) (string, error) {
	tranId, err := app.Setup.CreateUser(userId)
	if err != nil {
		return "", err
	}
	return tranId, nil
}

func (app *Application) GetUserBalance(userId string) (string, error) {
	balanceBytes, err := app.Setup.QueryBalanceByUserId(userId)
	if err != nil {
		return "", err
	}
	return string(balanceBytes), nil
}

func (app *Application) PushRequest(userId string, requirement string, reward float64, expiredTime string) (string, error) {
	u1, _ := uuid.NewV4()
	requestId := u1.String()
	createdTime := time.Now().Format(TIME_FORMAT)
	request := service.Request{
		RequestId:      requestId,
		Owner:          userId,
		Requirement:    requirement,
		Reward:         reward,
		Status:         0,
		CreateTime:     createdTime,
		ExpiredTime:    expiredTime,
		AcceptResponse: "",
		Responses:      nil,
	}
	tranId, err := app.Setup.PushRequest(request)
	if err != nil {
		return "", err
	}
	return tranId, nil
}

func (app *Application) GetAllRequests() ([]service.Request, error) {
	response, err := app.Setup.QueryAllRequest()
	if err != nil {
		return nil, err
	}
	var allRequests []service.Request
	err = json.Unmarshal([]byte(response), &allRequests)
	if err != nil {
		return nil, err
	}
	return allRequests, nil
}
func (app *Application) GetRequestByUserId(userId string) ([]service.Request, error) {
	response, err := app.Setup.QueryRequestByUserId(userId)
	if err != nil {
		return nil, err
	}
	var requests []service.Request
	err = json.Unmarshal([]byte(response), &requests)
	if err != nil {
		return nil, err
	}
	return requests, nil
}
func (app *Application) PushResponse(requestId string, owner string, answer string) (string, error) {
	u1, _ := uuid.NewV4()
	responseId := u1.String()
	createdTime := time.Now().Format(TIME_FORMAT)
	response := service.Response{
		RequestId:  requestId,
		ResponseId: responseId,
		Owner:      owner,
		Answer:     answer,
		CreateTime: createdTime,
	}
	tranId, err := app.Setup.PushRespone(response)
	if err != nil {
		return "", err
	}
	return tranId, nil
}
func (app *Application) GetResponseByUserId(userId string) ([]service.Response, error) {
	response, err := app.Setup.QueryResponseByUserId(userId)
	if err != nil {
		return nil, err
	}
	var responses []service.Response
	err = json.Unmarshal([]byte(response), &responses)
	if err != nil {
		return nil, err
	}
	return responses, nil
}
func (app *Application) GetResponseByRequestId(requestId string) ([]service.Response, error) {
	response, err := app.Setup.QueryResponseByRequestId(requestId)
	if err != nil {
		return nil, err
	}
	var responses []service.Response
	err = json.Unmarshal([]byte(response), &responses)
	if err != nil {
		return nil, err
	}
	return responses, nil
}
func (app *Application) AcceptResponse(userId string, requestId string, responseId string) (string, error) {
	tranId, err := app.Setup.AcceptResponse(userId, requestId, responseId)
	if err != nil {
		return "", err
	}
	return tranId, nil
}
