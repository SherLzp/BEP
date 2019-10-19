package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/BEP/bep_backend/models"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	BaseController
}

func (this *UserController) SignUp() {
	o := orm.NewOrm()
	var user models.User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		this.Error(UNKNOWN, "request body parse error", err)
		return
	}
	id, err := o.Insert(&user)
	if err != nil {
		this.Error(UNKNOWN, "create user error", err)
		return
	}
	this.Success(id, "user created")
}

func (this *UserController) SignIn() {
	o := orm.NewOrm()
	var user models.User
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	if err != nil {
		this.Error(UNKNOWN, "request body parse error", err)
		return
	}

	u := models.User{Email: user.Email}

	err = o.Read(&u, "Email")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		this.Error(UNKNOWN, "no such user", err)
		return
	}
	if u.Password != user.Password {
		this.Error(UNKNOWN, "password error", fmt.Errorf("password error"))
		return
	}

	this.Ctx.SetCookie("email", u.Email)
	this.Ctx.SetCookie("username", u.Username)
	this.Ctx.SetCookie("password", u.Password)

	this.Success(user, "login success")
}

const (
	pre_url string = "http://127.0.0.1:5000"
)

func (this *UserController) GenerateKeypair() {
	// call http://127.0.0.1:5000/generateKey
	//url := pre_url + "/generateKey"
	email := this.Ctx.GetCookie("email")
	password := this.Ctx.GetCookie("password")
	username := this.Ctx.GetCookie("username")

	user := models.User{Email: email, Username: username, Password: password}
	fmt.Println(user)

	// generate key 

	// save keypair to database

	this.Success(user, "generate keypair success")
}
