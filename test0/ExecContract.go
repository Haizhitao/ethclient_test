package test0

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Haizhitao/ethclient_test/store"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strings"
)

const contractAddress = "0x304Ed073eb82ae5C0464eD4c87Fc65910Df8c5Be"

func ExecContractByGo(rpcUrl string, _privateKey string) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		log.Fatal(err)
	}
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("age"))
	copy(value[:], []byte("30"))

	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash: ", tx.Hash().Hex())

	_, err = waitForReceipt(client, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	valueInContract, err := storeContract.Items(&bind.CallOpts{Context: context.Background()}, key)
	if err != nil {
		return
	}
	fmt.Println("is equal: ", valueInContract == value)
}

func ExecContractByAbi(rpcUrl string, _privateKey string) {
	abiBytes, err := os.ReadFile("Store_sol_Store.abi")
	if err != nil {
		log.Fatal(err)
	}
	abiString := string(abiBytes)

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 估算 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 准备交易数据
	contractABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal(err)
	}

	methodName := "setItem"
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("test_key"))
	copy(value[:], []byte("test_value"))
	input, err := contractABI.Pack(methodName, key, value)

	// 创建交易并签名交易
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 300000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	_, err = waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// 查询刚刚设置的值
	callInput, err := contractABI.Pack("items", key)
	if err != nil {
		log.Fatal(err)
	}
	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}

	// 解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte
	err1 := contractABI.UnpackIntoInterface(&unpacked, "items", result)
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Println("is equal: ", unpacked == value)

}

//Transaction sent: 0x482454abe4587c0027d0676e1b4a96503a3eabc84ed883d9ccf012a214c19bda
//is equal:  true
