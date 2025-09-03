package test0

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func TransactionInfo() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/e84d1cf5586b4f29ad9e236cfe5d2f89")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(9079246)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(tx.Value().String())    // 100000000000000000
		fmt.Println(tx.Gas())               // 21000
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		fmt.Println(tx.Nonce())             // 245132
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		txType := tx.Type()
		fmt.Printf("当前交易类型: %d\n", txType)

		// 1. 根据交易类型选择对应签名器
		var signer types.Signer
		switch tx.Type() {
		case types.LegacyTxType: // 类型 0：Legacy 交易
			signer = types.NewEIP155Signer(chainID)
		case types.DynamicFeeTxType: // 类型 2：EIP-1559 交易
			signer = types.NewLondonSigner(chainID)
		case types.BlobTxType: // 类型 3：EIP-4844 Blob 交易（需 go-ethereum 版本 >= v1.13.0）
			signer = types.NewCancunSigner(chainID)
		default:
			fmt.Printf("不支持的交易类型: %d\n", tx.Type())
			return
		}

		if sender, err := types.Sender(signer, tx); err == nil {
			fmt.Println("sender: ", sender.Hex())
		} else {
			log.Fatal(err)
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)   // []
		break
	}

	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal("TransactionCount: ", err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	txHash := common.HexToHash("0x2633811389812c3f3935ed9d7a4ff31133582d36a34781457369d0daccdcc090")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex())

	if sender, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
		fmt.Println("sender: ", sender.Hex())
	} else {
		log.Fatal(err)
	}

	fmt.Println("to: ", tx.To().Hex())
}
