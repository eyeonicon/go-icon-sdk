package main

import (
	"github.com/icon-project/goloop/client"
	"fmt"
	"paulrouge/go-icon-sdk/networks"
	"paulrouge/go-icon-sdk/wallet"
	// "github.com/icon-project/goloop/server/jsonrpc"
	"paulrouge/go-icon-sdk/transactions"
	"paulrouge/go-icon-sdk/util"

)


func main() {
	fmt.Println("Hello, world!")
	
	// set the active network globally (this way we can reuse the network id in the tx builders)
	networks.SetActiveNetwork(networks.Lisbon())
	

	Client := client.NewClientV3(networks.GetActiveNetwork().URL)
	_ = Client

	Wallet := wallet.LoadWallet("../mywallets/keystore.json", "joejoe")


	bn := util.ICXToLoop(1.5)

	// fmt.Println(bn)

	// fmt.Printf("type of bn: %T\n", bn)
		
	txobject := transactions.TransferICXBuilder("hx9c13cd371aed69c79870b3a3f7492c10122f0315", bn)

	tx, err := Client.SendTransaction(Wallet, txobject)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tx)



	// _ = txobject
}