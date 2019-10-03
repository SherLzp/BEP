package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
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
	RequestId  string `json:"request_id"`  // response to which
	ResponseId string `json:"response_id"` // unique response id
	Owner      string `json:"user_id"`     // response creator
	Answer     string `json:"answer"`      // answer(now it is a url)
	CreateTime string `json:"create_time"` // response create time
}

type User struct {
	UserId  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("fail to register chaincode event: %s", err)
	}
	fmt.Println("register chaincode event successfully: ", len(notifier))
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	fmt.Println("get event id: ", eventID)
	fmt.Println("get notifier: ", len(notifier))
	select {
	case ccEvent := <-notifier:
		fmt.Printf("receive chaincode event: %v\n", ccEvent)
	case <-time.After(time.Second * 60):
		return fmt.Errorf("cannot receive special chaincode event by event id(%s)", eventID)
	}
	return nil
}
