package main

import (
	"fmt"
	"github.com/BEP/bep_backend/fabric-sdk-go"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := fabric_sdk.FabricSetup{
		// Network parameters
		OrdererID: "orderer.BEP.com",

		// Channel parameters
		ChannelID:     "bepchannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/BEP/bep_backend/fixture/channel-artifacts/channel.tx",

		// Chaincode parameters
		ChainCodeID:     "bep",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/BEP/chaincode/",

		OrgAdmin:        "Admin",
		OrgName:         "OrgAlibaba",
		ConfigFile:      "/home/lxs/Application/go/src/github.com/BEP/config.yaml",

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
