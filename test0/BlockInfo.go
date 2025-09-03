package test0

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func Hello() {
	fmt.Println("hello world")
}

func BlockInfo() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/e84d1cf5586b4f29ad9e236cfe5d2f89")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(9074202)
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println(header.Number.Uint64())     // 5671744
	fmt.Println(header.Time)                // 1712798400
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1712798400
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println(len(block.Transactions()))   // 70
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count) // 70
	// 假设已通过 client.BlockByNumber 获取完整区块 block
	txs := block.Transactions()
	for _, tx := range txs {
		// 打印每笔交易的哈希
		fmt.Println("交易哈希:", tx.Hash().Hex())
		// 进一步获取交易发送者、金额等信息（需结合其他方法处理）
	}
}
