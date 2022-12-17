package main

import (
	"fmt"
	"github.com/icon-project/goloop/client"
	"github.com/eyeonicon/go-icon-sdk/networks"
	"github.com/eyeonicon/go-icon-sdk/transactions"
	"github.com/eyeonicon/go-icon-sdk/util"
	"github.com/eyeonicon/go-icon-sdk/wallet"
)


func main() {
	fmt.Println("Connecting to network...")
	
	// set the active network globally (this way we can reuse the network id in the tx builders)
	networks.SetActiveNetwork(networks.Lisbon())
	
	Client := client.NewClientV3(networks.GetActiveNetwork().URL)

	Wallet := wallet.LoadWallet("../mywallets/keystore01", "password")
	_ = Wallet

	// set the contract address
	contractAddress := "cx2b60e6e094df34a0d7c05b5ff5cb6758aba7e83e"
	
	// this address has a method called name that returns the current "name" value of the contract
	method := "name"

	// we only read the contract, so we don't need to sign the tx and can use the CallBuilder
	callObject := transactions.CallBuilder(contractAddress, method, nil)
	
	// send the call
	res, _ := Client.Call(callObject)

	fmt.Println(res) // Returns the current value of 'name' on the contract.

	// We will now try to change the value of 'name' on the contract.
	method = "setName"
	
	// the params for the method
	params := map[string]interface{}{
		"name": "Satoshi",
	}
	
	value := util.HexToBigInt("0x0")
	
	// We need to sign the tx, so we use the TransactionBuilder. We don't need to specify a value, so we pass in nil.
	tx := transactions.TransactionBuilder(Wallet.Address(), contractAddress, method, params, value)

	// sign the tx
	hash, err := Client.SendTransaction(Wallet, tx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*hash) // Returns the hash of the tx.

}