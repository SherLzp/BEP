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

	request := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "PushRespone", Args: [][]byte{b, []byte(eventID)}}
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

func (t *ServiceSetup) AcceptResponse(userId string, requestId Request, responseid Response) (string, error) {

	eventID := "eventAcceptResponse"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	b1 := []byte(userId)

	b2, err := json.Marshal(requestId)
	if err != nil {
		return "", fmt.Errorf("指定的request对象序列化时发生错误")
	}

	b3, err := json.Marshal(responseid)
	if err != nil {
		return "", fmt.Errorf("指定的response对象序列化时发生错误")
	}

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

func (t *ServiceSetup) CreateUser(usrid string) (string, error) {

	eventID := "eventCreateUser"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将usr对象序列化成为字节数组
	b, err := json.Marshal(usrid)
	if err != nil {
		return "", fmt.Errorf("error when serialize User")
	}

	request := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "CreateUser", Args: [][]byte{b, []byte(eventID)}}
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

func (t *ServiceSetup) QueryAllRequest() ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryAllRequest"}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) QueryRequestByUserId(userId string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryRequestByUserId", Args: [][]byte{[]byte(userId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) QueryResponseByUserId(userId string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryResponseByUserId", Args: [][]byte{[]byte(userId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) QueryBalanceByUserId(userId string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryBalanceByUserId", Args: [][]byte{[]byte(userId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) QueryResponseByRequestId(reqId string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryResponseByRequestId", Args: [][]byte{[]byte(reqId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}
