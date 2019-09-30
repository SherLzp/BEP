cryptogen generate --config=./crypto-config.yaml

configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channelartifacts/channel.tx -channelID bepchannel

configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/OrgAlibabaMSPanchors.tx -channelID bepchannel -asOrg OrgAlibabaMSP

configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/OrgBaiduMSPanchors.tx -channelID bepchannel -asOrg OrgBaiduMSP

