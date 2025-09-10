package test0

import (
	"context"
	"fmt"
	"github.com/Haizhitao/ethclient_test/store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
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
	//privateKey, err := crypto.HexToECDSA(_privateKey)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	//if err != nil {
	//	log.Fatal(err)
	//}

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("name"))
	copy(value[:], []byte("liuhaitao"))

	//tx, err := storeContract.SetItem(opt, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("tx hash: ", tx.Hash().Hex())
	valueInContract, err := storeContract.Items(&bind.CallOpts{Context: context.Background()}, key)
	if err != nil {
		return
	}
	fmt.Println("is equal: ", valueInContract == value)
}

func ExecContractByAbi() {

}
