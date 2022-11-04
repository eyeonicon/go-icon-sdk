package main

import (
	"github.com/icon-project/goloop/client"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	Client := client.NewClientV3("https://lisbon.net.solidwallet.io/api/v3")
	_ = Client
}