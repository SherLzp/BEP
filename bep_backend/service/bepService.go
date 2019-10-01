package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) PushReq(req Request) (string, error) {

	eventID := "eventAddReq"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将req对象序列化成为字节数组
	b, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("指定的req对象序列化时发生错误")
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

func (t *ServiceSetup) PushRes(res Response) (string, error) {

	eventID := "eventAddRes"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将res对象序列化成为字节数组
	b, err := json.Marshal(res)
	if err != nil {
		return "", fmt.Errorf("指定的res对象序列化时发生错误")
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

func (t *ServiceSetup) FindAllRequest(certNo, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryAllRequest", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindRequestByUserId(certNo, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryRequestByUserId", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindResponseByUserId(certNo, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryResponseByUserId", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindBalanceByUserId(certNo, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryBalanceByUserId", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindResponseByRequestId(certNo, name string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "QueryResponseByRequestId", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}