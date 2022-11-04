package main

import (
	"github.com/icon-project/goloop/client"
	"fmt"
	"paulrouge/go-icon-sdk/networks"
	"paulrouge/go-icon-sdk/transactions"

)


func main() {
	fmt.Println("Hello, world!")
	
	// set the active network (id) globally 
	networks.SetActiveNetwork(networks.Lisbon())
	
	Client := client.NewClientV3(networks.GetActiveNetwork().URL)
	_ = Client

	txobject := transactions.TransferICXBuilder("hx9c13cd371aed69c79870b3a3f7492c10122f0315", "1000000000000000000")
}