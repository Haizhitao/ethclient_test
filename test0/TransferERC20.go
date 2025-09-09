package test0

//课件地址 https://github.com/MetaNodeAcademy/Advanced1-backend-upgrade/blob/main/ethclient%E5%AE%9E%E6%88%98/2.06%20%E4%BB%A3%E5%B8%81%E8%BD%AC%E8%B4%A6/%E4%BB%A3%E5%B8%81%E8%BD%AC%E8%B4%A6.md
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func TransferERC20(rpcUrl string, _privateKey string) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return
	}
	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress:", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x462172A021d0b046f4c0e8A4F83Cc2255AE178fA")
	tokenAddress := common.HexToAddress("0xAC333bCf56D957b6e051fa278D7807d143BA5F46")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		//To:   &toAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

//操作过程：
//1. 铸造代币
//利用 remix，在sepolia网络，用账户 0xc76341B25002d55ff0B4AbD9FB1fEc974bf639A5， 部署了 RCCDemoToken 代币合约
//合约地址：0xAC333bCf56D957b6e051fa278D7807d143BA5F46
//用账户 0xc76341B25002d55ff0B4AbD9FB1fEc974bf639A5，通过metamsk 向合约发送了 0.002ETH，为本账户铸造了数量为 200,000 的代币（ERC-20: RCCDemoToken (RDT)）
//交易地址是 0x07d6eced9ce9554e2c34e524903ee701064c3844a3c6271a2159c74104a83c5e

//2. 代币转账
//替换该脚本中的合约地址为 0xAC333bCf56D957b6e051fa278D7807d143BA5F46
//替换toAddress为另一个账户地址：0x462172A021d0b046f4c0e8A4F83Cc2255AE178fA
//执行当前go脚本
//输出的交易地址为：0xde130313a48a7a48f5544eb97dea91e5927c9b6187301202fd5edf291895cce0
