package handler

import (
	"github.com/BEP/bep_backend/service"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"time"
)

const (
	TIME_FORMAT = "2018-02-02 02:02:02"
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

func (app *Application) PushRequest(userId string, requirement string, reward float64, duration float64) (string, error) {
	u1, _ := uuid.NewV4()
	requestId := u1.String()
	now := time.Now()
	endDuration, _ := time.ParseDuration(string(int32(duration * 24)))
	createdTime := now.Format(TIME_FORMAT)
	expiredTime := now.Add(endDuration).Format(TIME_FORMAT)
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
