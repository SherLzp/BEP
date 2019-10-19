package models

import (
	"fmt"
	"github.com/astaxie/beego"
)

type User struct {
	Id int64 `orm:"pk;auto" json:"id"`
	Email string `orm:"unique" json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Keypairs []*Keypair `orm:"reverse(many)" json:"keypairs"`
}

type Keypair struct {
	Id int64 `orm:"pk;auto" json:"id"`
	Pubkey string `orm:"unique" json:"public_key"`
	Privkey string `json:"private_key"`
	User *User `orm:"rel(fk)" json:"user"`
}

func NewUser(email string, username string, password string) (*User, error) {
	if username == "" || email == "" || password == "" {
		beego.Error("username/email/password empty")
		return nil, fmt.Errorf("username/email/password empty")
	}
	return &User{Email: email, Username: username, Password: password}, nil
}
