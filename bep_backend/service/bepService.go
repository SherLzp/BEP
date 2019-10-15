package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) PushRequest(req Request) (string, error) {

	eventID := "eventPushRequest"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将req对象序列化成为字节数组
	b, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("error when serialize Request")
	}

	request := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "PushRequest", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(request)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) PushRespone(res Response) (string, error) {

	eventID := "eventPushRespone"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将res对象序列化成为字节数组
	b, err := json.Marshal(res)
	if err != nil {
		return "", fmt.Errorf("error when serialize Response")
	}

	request := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "PushResponse", Args: [][]byte{b, []byte(eventID)}}
	response, err := t.Client.Execute(request)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(response.TransactionID), nil
}

func (t *ServiceSetup) AcceptResponse(userId string, requestId string, responseId string) (string, error) {

	eventID := "eventAcceptResponse"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	b1 := []byte(userId)
	b2 := []byte(requestId)
	b3 := []byte(responseId)

	request := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "AcceptResponse", Args: [][]byte{b1, b2, b3, []byte(eventID)}}
	respone, err := t.Client.Execute(request)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) CreateUser(userId string) (string, error) {
	eventID := "eventCreateUser"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	request := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "CreateUser", Args: [][]byte{[]byte(userId), []byte(eventID)}}
	respone, err := t.Client.Execute(request)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) QueryAllRequest() (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryAllRequest"}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) QueryRequestByUserId(userId string) (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryRequestByUserId", Args: [][]byte{[]byte(userId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", nil
	}
	return string(respone.Payload), nil
}

func (t *ServiceSetup) QueryResponseByUserId(userId string) (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryResponseByUserId", Args: [][]byte{[]byte(userId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) QueryBalanceByUserId(userId string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryBalanceByUserId", Args: [][]byte{[]byte(userId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) QueryResponseByRequestId(reqId string) (string, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryResponseByRequestId", Args: [][]byte{[]byte(reqId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) QueryRequestByRequestId(reqId string) (string, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryRequestByRequestId", Args: [][]byte{[]byte(reqId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) QueryResponseByResponseId(resId string) (string, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryResponseByResponseId", Args: [][]byte{[]byte(resId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}
