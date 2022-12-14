// Package networks provides a simple way to set the active network globally. 
// The url of the active network is the url of the node you want to connect to.
// 
// Set the active network with SetActiveNetwork(network Network) and 
// get the active network with GetActiveNetwork() Network.
// 
// The Network struct contains the URL and NID of the network.
// You can use the provided functions to get the URL and NID of the mainnet and the Lisbon testnet.
// And you can create your own Network by creating a function that returns a Network struct.

package networks

import (
	"github.com/icon-project/goloop/server/jsonrpc"
)

var activeNetwork Network

type Network struct {
	URL string
	NID jsonrpc.HexInt
}

// Network for Lisbon
func Lisbon() Network {
	return Network{
		URL: "https://lisbon.net.solidwallet.io/api/v3",
		NID: "0x2",
	}
}

// Network for Main net
func Mainnet() Network {
	return Network{
		URL: "https://ctz.solidwallet.io/api/v3",
		NID: "0x1",
	}
}

// Set the active network globally
func SetActiveNetwork(network Network) {
	activeNetwork = network
}

// Get the active network
func GetActiveNetwork() Network {
	return activeNetwork
}