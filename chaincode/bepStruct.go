package main

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
