package main

import (
	"fmt"
	pd "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type BepChaincode struct{}

// ********** chaincode begin ********** //
func (t *BepChaincode) Init(stub shim.ChaincodeStubInterface) pd.Response {
	return shim.Success(nil)
}

func (t *BepChaincode) Invoke(stub shim.ChaincodeStubInterface) pd.Response {

	fn, _ := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + fn)

	fmt.Println("invoke did not find func: " + fn)
	return shim.Error("Received unknown function invocation")
}

