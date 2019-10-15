package main

import (
	"fmt"
	"github.com/BEP/bep_backend/handler"
	_ "github.com/BEP/bep_backend/routers"
	"github.com/BEP/bep_backend/service"
	"github.com/BEP/sdkInit"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
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

	//-------------------------------------test start------------------------------------------------

	// create user Sher
	msg, err := serviceSetup.CreateUser("Sher")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add User successfully, transaction id is: ", msg)
	}

	// create user Jack
	msg, err = serviceSetup.CreateUser("Jack")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add User successfully, transaction id is: ", msg)
	}

	// create a request
	request1 := service.Request{
		RequestId:      "Request_001",
		Owner:          "Sher",
		Requirement:    "I want Jack's passport",
		Reward:         1.0,
		Status:         0,
		CreateTime:     "2019-10-01 10:00:00",
		ExpiredTime:    "2019-10-12 10:00:00",
		AcceptResponse: "",
		Responses:      nil,
	}

	// create a request
	request2 := service.Request{
		RequestId:      "Request_002",
		Owner:          "Sher",
		Requirement:    "This is a test",
		Reward:         1.0,
		Status:         0,
		CreateTime:     "2019-10-01 10:00:00",
		ExpiredTime:    "2019-10-12 10:00:00",
		AcceptResponse: "",
		Responses:      nil,
	}

	// create a response
	response1 := service.Response{
		RequestId:  "Request_001",
		ResponseId: "Response_001",
		Owner:      "Jack",
		Answer:     "It's in /home/machine/passport/Sher",
		CreateTime: "2019-10-11 10:00:00",
	}

	// push request
	msg, err = serviceSetup.PushRequest(request1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add Request successfully, transaction id is: ", msg)
	}

	// push request
	msg, err = serviceSetup.PushRequest(request2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add Request successfully, transaction id is: ", msg)
	}

	// push response
	msg, err = serviceSetup.PushRespone(response1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add Respone successfully, transaction id is: ", msg)
	}

	// accept response
	msg, err = serviceSetup.AcceptResponse("Sher", request1.RequestId, response1.ResponseId)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("AcceptResponse successfully, transaction id is: ", msg)
	}

	//-------------------------------------test over----------------------------------------------
	handler.App = handler.Application{
		Setup: &serviceSetup,
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	// start the server
	beego.Run()
}
