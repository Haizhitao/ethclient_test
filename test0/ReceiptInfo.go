package test0

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func ReceiptInfo(rpcUrl string) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(9118465)
	blockHash := common.HexToHash("0xd9279d01686204b8096b87422b6e37c7c48df0309c5e607ab1986b4bcdb3099b")

	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("receiptByHash[0]: ", receiptByHash[0])
	fmt.Println("receiptsByNum[0]: ", receiptsByNum[0])
	fmt.Println("pointer comparison: ", receiptByHash[0] == receiptsByNum[0]) //指针地址的比较
	//内容比较
	fmt.Println("content comparison: ", reflect.DeepEqual(receiptByHash[0], receiptsByNum[0]))

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.TransactionIndex)
		fmt.Println(receipt.ContractAddress.Hex())
		break
	}

	txHash := common.HexToHash("0x2633811389812c3f3935ed9d7a4ff31133582d36a34781457369d0daccdcc090")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)
	fmt.Println(receipt.TxHash.Hex())
	fmt.Println(receipt.TransactionIndex)
	fmt.Println(receipt.ContractAddress.Hex())
}
