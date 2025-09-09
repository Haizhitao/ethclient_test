package test0

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func Transfer(rpcUrl string, _privateKey string) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal("client.dial err: ", err)
	}
	privateKey, err := crypto.HexToECDSA(_privateKey) //账户私钥
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("can`t assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(10000000000000000) //0.01 ETH
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x462172A021d0b046f4c0e8A4F83Cc2255AE178fA")
	//var data []byte
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     nil,
	})

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx send: %s\n", signTx.Hash().Hex()) //0x04d94231281a7775f829cad4e82a7e5c2a7544a836e03c5640ceeeaa01639129

}
