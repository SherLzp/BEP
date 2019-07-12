package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pd "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type Request struct {
	RequestId      string     `json:"request_id"`   // unique request id
	Owner          string     `json:"user_id"`      // creator
	Requirement    string     `json:"requirement"`  // your requirement
	Reward         float64    `json:"reward"`       // the value of this question
	Status         int        `json:"status"`       // 0-under going 1-success 2-failure
	CreateTime     string     `json:"create_time"`  // request create time
	ExpiredTime    string     `json:"expired_time"` // expire time
	AcceptResponse string     `json:"accept_response_id"`
	Responses      []Response `json:"responses"`
}

type Response struct {
	RequestId  string `json:"request_id"`  // match which request
	ResponseId string `json:"response_id"` // unique response id
	Owner      string `json:"user_id"`     // response creator
	Answer     string `json:"answer"`      // answer(now it is a url)
	CreateTime string `json:"create_time"` // response create time
}

type User struct {
	UserId  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

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

func (t *BepChaincode) PushRequest(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6(request_id, user_id, requirement, reward, create_time, expired_time)")
	}
	requestId := args[0]
	userId := args[1]
	requirement := args[2]
	reward, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	status := 0
	createTime := args[4]
	expiredTime := args[5]
	request := Request{
		RequestId:      requestId,
		Owner:          userId,
		Requirement:    requirement,
		Reward:         reward,
		Status:         status,
		CreateTime:     createTime,
		ExpiredTime:    expiredTime,
		AcceptResponse: nil,
		Responses:      nil,
	}
	requestKey, err := stub.CreateCompositeKey("Request", []string{"request", requestId})
	if err != nil {
		return shim.Error(err.Error())
	}
	userRequestKey, err := stub.CreateCompositeKey("User_Request", []string{userId, requestId})
	if err != nil {
		return shim.Error(err.Error())
	}
	requestJSONasBytes, err := json.Marshal(request)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(requestKey, requestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(userRequestKey, requestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(requestId, requestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(requestJSONasBytes)
}

func (t *BepChaincode) PushResponse(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5(request_id, response_id, user_id, answer, create_time)")
	}
	requestId := args[0]
	// check if the request exists
	requestAsBytes, err := stub.GetState(requestId)
	if err != nil {
		return shim.Error("The request does not exist")
	}
	responseId := args[1]
	userId := args[2]
	// check the user is not the owner
	request := Request{}
	err = json.Unmarshal(requestAsBytes, &request)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check user
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

func (t *BepChaincode) AcceptResponse(stub shim.ChaincodeStubInterface, args []string) pd.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3(user_id, request_id, response_id)")
	}
	userId := args[0]
	requestId := args[1]
	responseId := args[2]
	// check the user is the owner of the request
	requestAsBytes, err := stub.GetState(requestId)
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
	_, err = stub.GetState(responseId)
	if err != nil {
		return shim.Error(err.Error())
	}
	// if everything is ok, then change state
	// change the request state, and rewrite ledger
	request.AcceptResponse = responseId
	request.Status = 1
	requestAsBytes, err = json.Marshal(request)
	if err != nil {
		return shim.Error(err.Error())
	}
	requestKey, err := stub.CreateCompositeKey("Request", []string{"request", requestId})
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
	err = stub.PutState(requestId, requestAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	// reward the user, and rewrite the ledger
	userAsBytes, err := stub.GetState(userId)
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
	err = stub.PutState(userId, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
