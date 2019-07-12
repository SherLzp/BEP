package main

import (
	"encoding/json"
	pd "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Request struct {
	RequestId string `json:"request_id"` // unique request id
	Owner string `json:"owner_id"` // creator
	Requirement string `json:"requirement"` // your requirement
	AcceptResponse int `json:"accept_response"`
	Reward float64 `json:"reward"` // the value of this question
	Status int `json:"status"` // 0-under going 1-success 2-failure
	CreateTime string `json:"create_time"` // request create time
	ExpiredTime string `json:"expired_time"`// expire time
	Responses []Response `json:"responses"`
}

type Response struct{
	RequestId string `json:"request_id"` // match which request
	ResponseId string `json:"response_id"` // unique response id
	Owner string `json:"owner_id"` // response creator
	Answer string `json:"answer"` // answer(now it is a url)
	CreateTIme string `json:"create_time"` // response create time
}

type User struct {
	UserId string `json:"user_id"`
	Balance float64 `json:"balance"`
}

func (t *BepChaincode) CreateUser(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 1(userid)")
	}
	userId := args[0]
	balance := 0.0
	user := User{userId, balance}

	// create a composite key, so that we can use prefix `user` to find all users
	userKey, err := stub.CreateCompositeKey("User", []string{"user", userId})
	if err != nil {
		return shim.Error(err.Error())
	}
	// convert struct to json
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}
	// save to ledger
	err = stub.PutPrivateData("collectionUser", userKey, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userJSONasBytes)
}

func (t *BepChaincode) PushRequest(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	
}

