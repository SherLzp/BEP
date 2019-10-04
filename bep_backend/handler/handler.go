package handler

import (
	"github.com/BEP/bep_backend/service"
	"github.com/astaxie/beego"
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
