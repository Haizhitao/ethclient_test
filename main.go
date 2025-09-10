package main

import (
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
	privateKey := getEnv("PRIVATE_KEY", "")
	wssRpcUrl := getEnv("WSS_RPC_URL", "")

	_ = privateKey
	_ = rpcUrl
	_ = wssRpcUrl

	//test0.BlockInfo()
	//test0.TransactionInfo()
	//test0.ReceiptInfo(rpcUrl)
	//test0.Wallet()
	//test0.Transfer(rpcUrl, privateKey)
	//test0.TransferERC20(rpcUrl, privateKey)
	//test0.GetBalance(rpcUrl)
	//test0.SubBlock(wssRpcUrl)
	//test0.DeployByAbigen(rpcUrl, privateKey)
	//test0.DeployByCode(rpcUrl, privateKey)
	//test0.LoadContract(rpcUrl)
	test0.ExecContractByGo(rpcUrl, privateKey)
}

func getEnv(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}
