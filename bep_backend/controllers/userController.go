package controllers

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
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
	email := this.Ctx.GetCookie("email")
	password := this.Ctx.GetCookie("password")
	username := this.Ctx.GetCookie("username")

	user := models.User{Email: email, Username: username, Password: password}
	fmt.Println(user)

	// generate key 
	url := pre_url + "/generateKey"
	res, err := http.Get(url)
	if err != nil {
		this.Error(UNKNOWN, "generate key error", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		this.Error(UNKNOWN, "read http response body error", err)
		return
	}

	var keypair models.Keypair
	err = json.Unmarshal(body, &keypair)
	if err != nil {
		this.Error(UNKNOWN, "parse json error", err)
		return
	}

	o := orm.NewOrm()
	err = o.Read(&user, "Email")
	if err != nil {
		this.Error(UNKNOWN, "user not found", err)
		return
	}

	keypair.User = &user
	_, err = o.Insert(&keypair)
	if err != nil {
		this.Error(UNKNOWN, "failed to save keypair", err)
		return
	}

	if o.Read(&user, "Email") == nil {
		user.Keypairs = append(user.Keypairs, &keypair)
		if _, err := o.Update(&user); err != nil {
			this.Error(UNKNOWN, "update user table error", err)
			return
		}
	} else {
		this.Error(UNKNOWN, "user not found", err)
		return
	}

	this.Success(0, "generate keypair success")
}
