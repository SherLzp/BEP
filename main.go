package main

import (
	"fmt"
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

}
