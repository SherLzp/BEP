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
