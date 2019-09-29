package main

import (
	"fmt"
	"./sdkInit"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := sdkInit.FabricSetup{
		// Network parameters
		OrdererID: "orderer.BEP.com",

		// Channel parameters
		ChannelID:     "bepchannel",
		// os.Getenv("GOPATH") +
		ChannelConfig: "/home/sher/go/src/github.com/BEP/bep_backend/fixture/channel-artifacts/channel.tx",

		// Chaincode parameters
		ChainCodeID:     "BEP",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/BEP/chaincode/",

		OrgAdmin:        "Admin",
		OrgName:         "OrgAlibaba",
		OrgID:			 "OrgAlibaba.BEP.com",
		ConfigFile:      "/home/sher/go/src/github.com/BEP/config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v \n", err)
		return
	}

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	// Close SDK
	defer fSetup.CloseSDK()
}
