package main

import (
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
		Requirement:    "I want sher's passport",
		Reward:         1.0,
		Status:         0,
		CreateTime:     "2019-10-01 10:00:00",
		ExpiredTime:    "2019-10-11 10:00:00",
		AcceptResponse: "",
		Responses:      nil,
	}

	request2 := service.Request{
		RequestId:      "Request_002",
		Owner:          "Sher",
		Requirement:    "I want sher's driver license",
		Reward:         10.0,
		Status:         0,
		CreateTime:     "2019-10-11 10:00:00",
		ExpiredTime:    "2019-10-21 10:00:00",
		AcceptResponse: "",
		Responses:      nil,
	}

	msg, err := serviceSetup.PushRequest(request1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add Request successfully, transaction id is: ", msg)
	}

	msg, err = serviceSetup.PushRequest(request2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Add Request successfully, transaction id is: ", msg)
	}
}
