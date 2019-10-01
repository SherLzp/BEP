cryptogen generate --config=./crypto-config.yaml
configtxgen -profile BEPOrdererGenesis -outputBlock ./channel-artifacts/genesis.block

configtxgen -profile BEPChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID bepchannel

configtxgen -profile BEPChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID bepchannel -asOrg Org1MSP

configtxgen -profile BEPChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID bepchannel -asOrg Org2MSP
