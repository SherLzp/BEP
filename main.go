package main

import (
	"fmt"
	"github.com/BEP/bep_backend/service"
	"github.com/BEP/sdkInit"
	"os"
)

const (
	configFile = "config.yaml"
	initialized = false
	SimpleCC = "simplecc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "bepchannel",
		ChannelConfig: "/home/sher/go/src/github.com/BEP/fixtures/channel-artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.bep.com",

		ChaincodeID: SimpleCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/BEP/chaincode/",
		UserName:"User1",

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

	// interact with fabric network
	serviceSetup := service.ServiceSetup{
		ChaincodeID: "bep",
		Client: channelClient,
	}

	req1 := service.Request{
		RequestId: "1",
		Owner: "p1",
		Requirement: "I want an apple",
		Reward: 5.12,
		Status: 0,
		CreateTime: "2019-10-01-16-35",
		ExpiredTime: "2019-10-07-16-35",
		AcceptResponse: "",
		Responses: nil,
	}

	req2 := service.Request{
		RequestId: "2",
		Owner: "p2",
		Requirement: "I want a banana",
		Reward: 10.24,
		Status: 0,
		CreateTime: "2019-10-01-16-35",
		ExpiredTime: "2019-10-07-16-35",
		AcceptResponse: "",
		Responses: nil,
	}

	msg, err := serviceSetup.PushReq(req1)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	msg, err = serviceSetup.PushReq(req2)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
}
