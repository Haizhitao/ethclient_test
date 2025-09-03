package test0

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func Wallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("privateKey: ", hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, OK := publicKey.(*ecdsa.PublicKey)
	if !OK {
		log.Fatal("assert type error")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKey: ", hexutil.Encode(publicKeyBytes)[4:])
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)

	//手动实现以太坊地址的哈希计算逻辑
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}

//privateKey:  453d85c9c0207ea9f440c3c8a1d1280981d6a4bacfae9aa4e3969652e3b1e559
//publicKey:  896137e6ea6ed88c01ba1cedd2c05f4a8459a5e16b5d9d3a31c161c2d4abefd3b663e3a3718f8d1f4d35fa55015edf5e5236e248d2f0f5cd92bc1b935b058ef3
//address:  0x7559767A85e0dd77D521baF51B6D27cC49F1A8ff
//full: 0x0f399a78c33128c7a19791e57559767a85e0dd77d521baf51b6d27cc49f1a8ff
//0x7559767a85e0dd77d521baf51b6d27cc49f1a8ff
