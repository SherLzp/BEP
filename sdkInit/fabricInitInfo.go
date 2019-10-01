package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

type InitInfo struct {
	ChannelID     string // 通道id
	ChannelConfig string // 通道配置文件
	OrgAdmin      string // Admin
	OrgName       string // 组织名称层
	OrdererOrgName	string // Orderer组织名称
	OrgResMgmt *resmgmt.Client

	ChaincodeID	string
	ChaincodeGoPath	string
	ChaincodePath	string
	UserName	string
}
