package models

import "github.com/BEP/bep_backend/service"

type YourResponse struct {
	Request  service.Request  `json:"request"`
	Response service.Response `json:"response"`
}

type ReceivedResponse struct {
	Request   service.Request    `json:"request"`
	Responses []service.Response `json:"responses"`
}
