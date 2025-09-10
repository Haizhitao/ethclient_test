package test0

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"strings"
)

func SubscribeEvent(rpcUrl string) {
	abiByets, err := os.ReadFile("StoreIndexed_sol_Store.abi")
	if err != nil {
		log.Fatal(err)
	}
	abiString := string(abiByets)
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0xc9ce9cd2df9a0fc226b17247d1255b679a1a6326") //使用 StoreIndexed.sol 在 remix 部署的合约
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}

}

//0x53c544d19c88b3d39f13bb9dc1e7c7b7c4a07aaba5a9a0305a6c1b117a6a79fd
//9174835
//0x0343712f55fff2570ebe5c8355e21b70b9b53f6dc75b52ecd247ce619e5b0744
//0000000000000000000000000000000000000000000000000000000000000000
//0000000000000000000000000000000000000000000000000000000000002200
//topics[0]= 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
//index topic: [0x0000000000000000000000000000000000000000000000000000000000002200]
