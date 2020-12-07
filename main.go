package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {
	//testCertContent := "-----BEGIN CERTIFICATE-----\nMIICKjCCAdCgAwIBAgIQNQ0eFR6ua5BuEudbydW+CDAKBggqhkjOPQQDAjBzMQsw\nCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy\nYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEcMBoGA1UEAxMTY2Eu\nb3JnMS5leGFtcGxlLmNvbTAeFw0yMDEyMDcwMjIzMDBaFw0zMDEyMDUwMjIzMDBa\nMGwxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1T\nYW4gRnJhbmNpc2NvMQ8wDQYDVQQLEwZjbGllbnQxHzAdBgNVBAMMFlVzZXIxQG9y\nZzEuZXhhbXBsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATfhRyr1/9b\n/Th+qiueJdAGvUl/BWeJESTxoE+Iuc9tp2G0RRc5LZ1ckuX0HDZvUqILzJ0dtmT4\n9Muq+UR2w7vmo00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADArBgNV\nHSMEJDAigCDiKrzVO/sn4SkUH+iRd74cVLNkP35LEOE6nXPhg3/VgzAKBggqhkjO\nPQQDAgNIADBFAiEAmbD1voYRa/hCUPFtLb/8Ds/FoB4ixF1AQfj0fN5C/08CIAnT\nbau9LQ73EanTTWx6tkcN/3lGusX8tFsXbLjcIDi6\n-----END CERTIFICATE-----\n"
	//testPrivKeyContent := "-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgNPRGDAoTKi2rEwHf\nMWl2jTvruZ2MO6KGsP8CkFmm3MahRANCAATfhRyr1/9b/Th+qiueJdAGvUl/BWeJ\nESTxoE+Iuc9tp2G0RRc5LZ1ckuX0HDZvUqILzJ0dtmT49Muq+UR2w7vm\n-----END PRIVATE KEY-----\n"
	wallet,err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Println("can't find wallet path")
	}
	//wallet.Put("user1", gateway.NewX509Identity("Org1MSP",testCertContent, testPrivKeyContent))

	gw, err := gateway.Connect(gateway.WithConfig(config.FromFile("config_e2e.yaml")),
		gateway.WithIdentity(wallet, "user1"))

	network, err := gw.GetNetwork("mychannel")

	if err != nil {
		fmt.Printf("Failed to get network: %s", err)
		return
	}

	contract := network.GetContract("mycc")

	result, err := contract.EvaluateTransaction("query","a")

	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s", err)
		return
	}
	fmt.Println(string(result))

	result, err = contract.SubmitTransaction("invoke", "a", "b", "2")

	if err != nil {
		fmt.Printf("Failed to submit transaction: %s", err)
		return
	}

	// the user might prefer a 'helper' method to reduce boilerplate error handling
	aValue,err := contract.EvaluateTransaction("query", "a")
	if err !=nil {
		fmt.Println("query failure")
	}

	fmt.Println(string(aValue))

}
