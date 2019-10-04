package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pd "github.com/hyperledger/fabric/protos/peer"
)

type BepChaincode struct{}

// ********** chaincode begin ********** //
func (t *BepChaincode) Init(stub shim.ChaincodeStubInterface) pd.Response {
	return shim.Success(nil)
}

func (t *BepChaincode) Invoke(stub shim.ChaincodeStubInterface) pd.Response {

	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running: " + fn)

	if fn == "CreateUser" {
		return t.CreateUser(stub, args)
	} else if fn == "PushRequest" {
		return t.PushRequest(stub, args)
	} else if fn == "PushResponse" {
		return t.PushResponse(stub, args)
	} else if fn == "AcceptResponse" {
		return t.AcceptResponse(stub, args)
	} else if fn == "QueryAllRequest" {
		return t.QueryAllRequest(stub)
	} else if fn == "QueryRequestByUserId" {
		return t.QueryRequestByUserId(stub, args)
	} else if fn == "QueryResponseByUserId" {
		return t.QueryResponseByUserId(stub, args)
	} else if fn == "QueryBalanceByUserId" {
		return t.QueryBalanceByUserId(stub, args)
	} else if fn == "QueryResponseByRequestId" {
		return t.QueryResponseByRequestId(stub, args)
	}

	fmt.Println("invoke did not find func: " + fn)
	return shim.Error("Received unknown function invocation")
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(BepChaincode))
	if err != nil {
		fmt.Printf("Error starting BEP Service chaincode: %s", err)
	}
}