package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func mockPushRequest(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("PushRequest"), []byte(args[0]), []byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5])})
	if res.Status != shim.OK {
		fmt.Println("PushRequest failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func TestQueryAllRequest(t *testing.T) {
	ast := assert.New(t)
	bep := new(BepChaincode)
	mstub := shim.NewMockStub("BepChaincode", bep)
	mockInit(t, mstub, nil)
	mockPushRequest(t, mstub, []string{"1", "p1", "I want an apple", "5.12", "2019-07-15-16-35", "2019-07-15-46-35"})
	mockPushRequest(t, mstub, []string{"2", "p2", "I want a banana", "10.24", "2019-07-16-16-35", "2019-07-16-46-35"})
	mockPushRequest(t, mstub, []string{"3", "p3", "I want a pineapple", "20.48", "2019-07-17-16-35", "2019-07-17-46-35"})

	res := mstub.MockInvoke("1", [][]byte{[]byte("QueryAllRequest")})
	ast.Equal(int(res.Status), shim.OK, "Fail to query all the requests.")
	ast.NotNil(res.Payload, "Fail to push the request, query finds nothing.")
	buffer := bytes.NewBuffer(res.Payload)
	req := Request{}

	cnt := 1
	for buffer.Len() > 0 {
		cur, err := buffer.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Fail to read the buffer.cnt = %d\n", cnt)
		}
		// if there's still a Request in the buffer
		if buffer.Len() > 0 {
			cur = cur[:len(cur)-1]
		}
		err = json.Unmarshal(cur, &req)
		ast.Nil(err, "Fail to unmarshal the request info.")

		switch cnt {
		case 1:
			ast.Equal("1", req.RequestId,  "RequestId not equal to 1.")
			ast.Equal("p1", req.Owner, "Owner ID should be p1.")
			ast.Equal(5.12, req.Reward, "reward should be 5.12.")
			ast.Equal("I want an apple", req.Requirement, "Requirement not equal to apple.")
			ast.Equal("2019-07-15-16-35", req.CreateTime, "CreateTime not equal.")
			ast.Equal("2019-07-15-46-35", req.ExpiredTime, "ExpiredTime not equal.")
		case 2:
			ast.Equal("2", req.RequestId,  "RequestId not equal to 2.")
			ast.Equal("p2", req.Owner, "Owner ID should be p2.")
			ast.Equal(10.24, req.Reward, "reward should be 10.24.")
			ast.Equal("I want a banana", req.Requirement, "Requirement not equal to banana.")
			ast.Equal("2019-07-16-16-35", req.CreateTime, "CreateTime not equal.")
			ast.Equal("2019-07-16-46-35", req.ExpiredTime, "ExpiredTime not equal.")
		case 3:
			ast.Equal("3", req.RequestId,  "RequestId not equal to 3.")
			ast.Equal("p3", req.Owner, "Owner ID should be p3.")
			ast.Equal(20.48, req.Reward, "reward should be 20.48.")
			ast.Equal("I want a pineapple", req.Requirement, "Requirement not equal to pineapple.")
			ast.Equal("2019-07-17-16-35", req.CreateTime, "CreateTime not equal.")
			ast.Equal("2019-07-17-46-35", req.ExpiredTime, "ExpiredTime not equal.")
		}
		cnt += 1
	}

}

func TestQueryRequestByUserId(t *testing.T) {
	ast := assert.New(t)
	bep := new(BepChaincode)
	mstub := shim.NewMockStub("BepChaincode", bep)
	mockInit(t, mstub, nil)
	mockPushRequest(t, mstub, []string{"1", "p1", "I want an apple", "5.12", "2019-07-15-16-35", "2019-07-15-46-35"})
	mockPushRequest(t, mstub, []string{"2", "p1", "I want a banana", "10.24", "2019-07-16-16-35", "2019-07-16-46-35"})
	mockPushRequest(t, mstub, []string{"3", "p1", "I want a pineapple", "20.48", "2019-07-17-16-35", "2019-07-17-46-35"})

	res := mstub.MockInvoke("1", [][]byte{[]byte("QueryRequestByUserId"), []byte("p1")})
	ast.Equal(int(res.Status), shim.OK, "Fail to query all the requests.")
	ast.NotNil(res.Payload, "query finds nothing!")
	buffer := bytes.NewBuffer(res.Payload)
	req := Request{}

	cnt := 1
	for buffer.Len() > 0 {
		cur, err := buffer.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Fail to read the buffer.cnt = %d\n", cnt)
		}
		// if there's still a Request in the buffer
		if buffer.Len() > 0 {
			cur = cur[:len(cur)-1]
		}
		err = json.Unmarshal(cur, &req)
		ast.Nil(err, "Fail to unmarshal the request info.")

		switch cnt {
		case 1:
			ast.Equal("1", req.RequestId,  "RequestId not equal to 1.")
			ast.Equal("p1", req.Owner, "Owner ID should be p1.")
			ast.Equal(5.12, req.Reward, "reward should be 5.12.")
			ast.Equal("I want an apple", req.Requirement, "Requirement not equal to apple.")
			ast.Equal("2019-07-15-16-35", req.CreateTime, "CreateTime not equal.")
			ast.Equal("2019-07-15-46-35", req.ExpiredTime, "ExpiredTime not equal.")
		case 2:
			ast.Equal("2", req.RequestId,  "RequestId not equal to 2.")
			ast.Equal("p1", req.Owner, "Owner ID should be p1.")
			ast.Equal(10.24, req.Reward, "reward should be 10.24.")
			ast.Equal("I want a banana", req.Requirement, "Requirement not equal to banana.")
			ast.Equal("2019-07-16-16-35", req.CreateTime, "CreateTime not equal.")
			ast.Equal("2019-07-16-46-35", req.ExpiredTime, "ExpiredTime not equal.")
		case 3:
			ast.Equal("3", req.RequestId,  "RequestId not equal to 3.")
			ast.Equal("p1", req.Owner, "Owner ID should be p1.")
			ast.Equal(20.48, req.Reward, "reward should be 20.48.")
			ast.Equal("I want a pineapple", req.Requirement, "Requirement not equal to pineapple.")
			ast.Equal("2019-07-17-16-35", req.CreateTime, "CreateTime not equal.")
			ast.Equal("2019-07-17-46-35", req.ExpiredTime, "ExpiredTime not equal.")
		}
		cnt += 1
	}
}
