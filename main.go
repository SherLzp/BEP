package main

import (
	"encoding/json"
	"fmt"
	"github.com/BEP/bep_backend/service"
	"github.com/BEP/sdkInit"
	"os"
)

const (
	configFile  = "config.yaml"
	initialized = false
	BepCC       = "bepcc"
)

func main() {
	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/BEP/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     BepCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/BEP/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	serviceSetup := service.ServiceSetup{
		ChaincodeID: BepCC,
		Client:      channelClient,
	}

	request1 := service.Request{
		RequestId:      "Request_001",
		Owner:          "Sher",
		Requirement:    "I want Sher's passport",
		Reward:         1.0,
		Status:         0,
		CreateTime:     "2019-10-01 10:00:00",
		ExpiredTime:    "2019-10-12 10:00:00",
		AcceptResponse: "",
		Responses:      nil,
	}

	response1 := service.Response{
		RequestId:      "Request_001",
		ResponseId:		"Response_001",
		Owner:          "Jack",
		Answer:			"It's in /home/machine/passport/Sher",
		CreateTime:     "2019-10-11 10:00:00",
	}

	msg, err := serviceSetup.CreateUser("Sher")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add User successfully, transaction id is: ", msg)
	}

	msg, err = serviceSetup.CreateUser("Jack")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add User successfully, transaction id is: ", msg)
	}

	msg, err = serviceSetup.PushRequest(request1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add Request successfully, transaction id is: ", msg)
	}

	result, err := serviceSetup.QueryRequestByUserId("Sher")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var request service.Request
		json.Unmarshal(result, &request)
		fmt.Println("根据userid查询request成功：")
		fmt.Println("This request belongs to:" + request.Owner + ", it's create time is:" + request.CreateTime)
	}

	// the 2 func below have bugs.
	msg, err = serviceSetup.PushRespone(response1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add Respone successfully, transaction id is: ", msg)
	}

	msg, err = serviceSetup.AcceptResponse("Sher", request1, response1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("AcceptResponse successfully, transaction id is: ", msg)
	}

}
