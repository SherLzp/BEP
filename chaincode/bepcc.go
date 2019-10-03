package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pd "github.com/hyperledger/fabric/protos/peer"
)

const (
	REQUEST_TYPE  = "requestObj"
	RESPONSE_TYPE = "responseObj"
)

/* Create a user */
func (t *BepChaincode) CreateUser(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1(userid)")
	}
	userId := args[0]
	balance := 0.0
	user := User{userId, balance}

	// convert struct to json
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}
	// save to ledger
	err = stub.PutState(userId, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userJSONasBytes)
}

func PutRequest(stub shim.ChaincodeStubInterface, request Request) ([]byte, bool) {
	requestJSONasBytes, err := json.Marshal(request)
	if err != nil {
		return nil, false
	}
	requestKey, err := stub.CreateCompositeKey("Request", []string{"request", request.RequestId})
	if err != nil {
		return nil, false
	}
	userRequestKey, err := stub.CreateCompositeKey("User_Request", []string{request.Owner, request.RequestId})
	//userRequestKey, err := stub.CreateCompositeKey(userId, []string{requestId})
	if err != nil {
		return nil, false
	}
	err = stub.PutState(requestKey, requestJSONasBytes)
	if err != nil {
		return nil, false
	}
	err = stub.PutState(userRequestKey, requestJSONasBytes)
	if err != nil {
		return nil, false
	}
	return nil, true
}

/* Push a request to the ledger */
func (t *BepChaincode) PushRequest(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2(request,event)")
	}
	var request Request
	err := json.Unmarshal([]byte(args[0]), &request)
	if err != nil {
		return shim.Error("error when deserialize Request")
	}
	_, status := PutRequest(stub, request)
	if !status {
		return shim.Error("error when PutRequest")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Push Request successfully"))
}

/* Push a response to a specific request */
func (t *BepChaincode) PushResponse(stub shim.ChaincodeStubInterface, args []string) pd.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5(request_id, response_id, user_id, answer, create_time)")
	}
	requestId := args[0]
	// check if the request exists
	requestKey, err := stub.CreateCompositeKey("Request", []string{"request", requestId})
	if err != nil {
		return shim.Error(err.Error())
	}
	requestAsBytes, err := stub.GetState(requestKey)
	if err != nil {
		return shim.Error("The request does not exist: " + err.Error())
	}
	responseId := args[1]
	userId := args[2]

	// check : whether this user push the request
	request := Request{}
	err = json.Unmarshal(requestAsBytes, &request)
	if err != nil {
		return shim.Error(err.Error())
	}
	if request.Owner == userId {
		return shim.Error("You cannot answer your own question")
	}

	answer := args[3]
	createTime := args[4]
	response := Response{
		RequestId:  requestId,
		ResponseId: responseId,
		Owner:      userId,
		Answer:     answer,
		CreateTime: createTime,
	}

	// add : put the response into the request
	request.Responses = append(request.Responses, response)
	requestAsBytes, err = json.Marshal(request)
	if err != nil {
		return shim.Error(err.Error())
	}
	userRequestKey, err := stub.CreateCompositeKey("User_Request", []string{userId, requestId})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(requestKey, requestAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(userRequestKey, requestAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// add : put the response into the ledger
	responseKey, err := stub.CreateCompositeKey("Response", []string{requestId, responseId})
	if err != nil {
		return shim.Error(err.Error())
	}
	userResponseKey, err := stub.CreateCompositeKey("User_Response", []string{userId, responseId})
	if err != nil {
		return shim.Error(err.Error())
	}
	responseJSONasBytes, err := json.Marshal(response)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(responseKey, responseJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(userResponseKey, responseJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(responseId, responseJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(responseJSONasBytes)
}

/* Push a accept to a specific response */
func (t *BepChaincode) AcceptResponse(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3(user_id, request_id, response_id)")
	}
	userId := args[0]
	requestId := args[1]
	responseId := args[2]
	// check whether the user is the owner of the request
	requestKey, err := stub.CreateCompositeKey("Request", []string{"request", requestId})
	if err != nil {
		return shim.Error(err.Error())
	}
	requestAsBytes, err := stub.GetState(requestKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	request := Request{}
	err = json.Unmarshal(requestAsBytes, &request)
	if err != nil {
		return shim.Error(err.Error())
	}
	if userId != request.Owner {
		return shim.Error("You are not the owner of the request")
	}

	// check the response exists
	responseAsBytes, err := stub.GetState(responseId)
	if err != nil {
		return shim.Error(err.Error())
	}

	// if everything is ok, then change the request state
	request.AcceptResponse = responseId
	request.Status = 1
	requestAsBytes, err = json.Marshal(request)
	if err != nil {
		return shim.Error(err.Error())
	}
	userRequestKey, err := stub.CreateCompositeKey("User_Request", []string{userId, requestId})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(requestKey, requestAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(userRequestKey, requestAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// reward the responser
	response := Response{}
	err = json.Unmarshal(responseAsBytes, &response)
	if err != nil {
		return shim.Error(err.Error())
	}
	userAsBytes, err := stub.GetState(response.Owner)
	if err != nil {
		return shim.Error(err.Error())
	}
	user := User{}
	err = json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}
	reward := request.Reward
	user.Balance = user.Balance + reward
	userAsBytes, err = json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(response.Owner, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

/* Query all requests in the ledger */
func (t *BepChaincode) QueryAllRequest(stub shim.ChaincodeStubInterface) pd.Response {
	resultIter, err := stub.GetStateByPartialCompositeKey("Request", []string{"request"})
	if err != nil {
		return shim.Error("Failed to query requests: " + err.Error())
	}
	defer resultIter.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultIter.HasNext() {
		queryResponse, err := resultIter.Next()
		if err != nil {
			return shim.Error("Fail to get request: " + err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteByte('\n')
		}

		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	return shim.Success(buffer.Bytes())
}

/* Query all requests pushed by UserId */
func (t *BepChaincode) QueryRequestByUserId(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1(userid)")
	}
	useridAsString := args[0]
	resultIter, err := stub.GetStateByPartialCompositeKey("User_Request", []string{useridAsString})
	//resultIter, err := stub.GetStateByPartialCompositeKey(useridAsString,[]string{""})
	if err != nil {
		return shim.Error("Failed to query requests: " + err.Error())
	}
	defer resultIter.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultIter.HasNext() {
		queryResponse, err := resultIter.Next()
		if err != nil {
			return shim.Error("Fail to get request: " + err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString("\n")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	return shim.Success(buffer.Bytes())
}

/* Query all responses pushed by UserId */
func (t *BepChaincode) QueryResponseByUserId(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1(userid)")
	}
	useridAsString := args[0]
	resultIter, err := stub.GetStateByPartialCompositeKey("User_Response", []string{useridAsString})
	if err != nil {
		return shim.Error("Failed to query response: " + err.Error())
	}
	defer resultIter.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultIter.HasNext() {
		queryResponse, err := resultIter.Next()
		fmt.Printf("user response: %s\n", queryResponse.Value)
		if err != nil {
			return shim.Error("Fail to get response: " + err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString("\n")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	return shim.Success(buffer.Bytes())
}

/* Query balancec by UserId */
func (t *BepChaincode) QueryBalanceByUserId(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1(userid)")
	}
	useridAsString := args[0]
	userAsBytes, err := stub.GetState(useridAsString)
	if err != nil {
		return shim.Error("Fail to get the user by id: " + err.Error())
	} else if userAsBytes == nil {
		return shim.Error("User does not exist")
	}

	usr := User{}
	err = json.Unmarshal(userAsBytes, &usr)
	if err != nil {
		return shim.Error("Fail to unmarshal the user: " + err.Error())
	}

	return shim.Success([]byte(fmt.Sprintf("%f", usr.Balance)))
}

/* Query all the responses by requestId */
func (t *BepChaincode) QueryResponseByRequestId(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1(userid)")
	}
	reqidAsString := args[0]
	requestKey, err := stub.CreateCompositeKey("Request", []string{"request", reqidAsString})
	if err != nil {
		return shim.Error(err.Error())
	}
	requestAsByte, err := stub.GetState(requestKey)

	request := &Request{}
	err = json.Unmarshal(requestAsByte, &request)
	if err != nil {
		return shim.Error(err.Error())
	}

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for _, curRes := range request.Responses {

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString("\n")
		}

		responseAsByte, err := json.Marshal(curRes)
		if err != nil {
			return shim.Error(err.Error())
		}

		buffer.WriteString(string(responseAsByte))
		bArrayMemberAlreadyWritten = true
	}

	return shim.Success(buffer.Bytes())
}
