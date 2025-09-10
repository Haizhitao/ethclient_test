package test0

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strings"
)

func QueryEvent(rpcUrl string) {
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
		FromBlock: big.NewInt(9171030),
		// ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
		// Topics: [][]common.Hash{
		//  {},
		//  {},
		// },
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal(err)
	}
	for _, vLog := range logs {
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
			fmt.Println("indexed topics:", topics[1:])
		}
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("signature topics=", hash.Hex())
}

//0xcb02281df520ff25814933caa15c50d907a88c1369f3eafd4dd9e222ac4080ba
//9173510
//0x6f419d9731180e45e61dbac9fe0ff15fbe0df17c00122d599ba24b5d33bee443
//0000000000000000000000000000000000000000000000000000000000000000
//0000000000000000000000000000000000000000000000000000000000002200
//topics[0]= 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
//indexed topics: [0x68656c6c6f000000000000000000000000000000000000000000000000000000]
//signature topics= 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
