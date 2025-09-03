package main

import (
	"fmt"
	"github.com/Haizhitao/ethclient_test/test0"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("未找到 .evn 文件")
	}
	rpcUrl := getEnv("RPC_URL", "")
	fmt.Println("rpcUrl: ", rpcUrl)

	//test0.BlockInfo()
	//test0.TransactionInfo()
	//test0.ReceiptInfo(rpcUrl)
	test0.Wallet()
}

func getEnv(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
