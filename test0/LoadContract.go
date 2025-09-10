package test0

import (
	"github.com/Haizhitao/ethclient_test/store2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

const (
	contractAddr = "0x304Ed073eb82ae5C0464eD4c87Fc65910Df8c5Be"
)

func LoadContract(rpcUrl string) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store2.NewStore2(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	_ = storeContract
}
