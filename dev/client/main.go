package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TestType string

const (
	abiFile = "./contracts/counter/counter.abi"
	binFile = "./contracts/counter/counter.bin"

	Oip20Test   = TestType("oip20")
	CounterTest = TestType("counter")
)

func main() {
	testTypeParam := flag.String("type", "oip20", "choose which test to run")
	flag.Parse()

	privKey := []string{
		"8d757e29cf7636d07d2f8ea2aaa14051fd173874ead2a0aa22e7ef4770734b4c",
		//"171786c73f805d257ceb07206d851eea30b3b41a2170ae55e1225e0ad516ef42",
		//"b7700998b973a2cae0cb8e8a328171399c043e57289735aca5f2419bd622297a",
		//"00dcf944648491b3a822d40bf212f359f699ed0dd5ce5a60f1da5e1142855949",
	}

	var testFunc func(privKey string, blockTime time.Duration) error
	switch TestType(*testTypeParam) {
	case Oip20Test:
		fmt.Printf("contract: %s\n", *testTypeParam)
		testFunc = standardOip20Test
		break
	default:
		fmt.Printf("contract: %s\n", CounterTest)
		testFunc = counterTest
	}

	for _, k := range privKey {
		test := func(key string) {
			testFunc(key, time.Millisecond*5000)
		}
		go writeRoutine(test, k)
	}
	<-make(chan struct{})
}

func writeRoutine(test func(string), key string) {
	for {
		test(key)
		log.Printf("recover writeRoutine...")
		sleep(3)
	}
}

func counterTest(privKey string, blockTime time.Duration) error {
	var (
		privateKey    *ecdsa.PrivateKey
		senderAddress common.Address
	)

	privateKey, senderAddress = initKey(privKey)
	counterContract := newContract("counter", "", abiFile, binFile)

	client, err := ethclient.Dial(RpcUrl)
	if err == nil {
		err = deployContract(client, senderAddress, privateKey, counterContract, 3)
	}

	for err == nil {
		err = writeContract(client, counterContract, senderAddress, privateKey, nil, blockTime, "add", big.NewInt(100))
		uint256Output(client, counterContract, "getCounter")
		err = writeContract(client, counterContract, senderAddress, privateKey, nil, blockTime, "subtract")
		uint256Output(client, counterContract, "getCounter")
	}
	return err
}

func standardOip20Test(privKey string, blockTime time.Duration) error {
	privateKey, sender := initKey(privKey)

	client, err := ethclient.Dial(RpcUrl)
	if err != nil {
		log.Printf("failed to dial: %+v", err)
	}

	oip20, auth, err := deployOip(client, sender, privateKey)
	if err != nil {
		log.Printf("failed to deploy: %+v", err)
	}

	toAddress := common.HexToAddress("0x78B63831Fb1050841DEaBE5cc785bCaA91AF3478")
	for err == nil {
		nonce, err := transferOip(client, oip20, sender, auth, toAddress)
		if err != nil {
			log.Printf("failed to transfer Oip: %+v", err)
			break
		}
		fmt.Printf(
			"==================================================\n"+
				"Standard OIP20 transfer:\n"+
				"	from					: <%s>\n"+
				"	nonce					: <%d>\n"+
				"	to					: <%s>\n",
			sender, nonce, toAddress,
		)
		time.Sleep(blockTime)
	}

	return err
}
