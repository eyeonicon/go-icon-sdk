package networks

import (
	"github.com/icon-project/goloop/server/jsonrpc"
)

var activeNetwork Network

type Network struct {
	URL string
	NID jsonrpc.HexInt
}

// make a Network for Lisbon
func Lisbon() Network {
	return Network{
		URL: "https://lisbon.net.solidwallet.io/api/v3",
		NID: "0x2",
	}
}

// make a Network for Main net
func Mainnet() Network {
	return Network{
		URL: "https://ctz.solidwallet.io/api/v3",
		NID: "0x1",
	}
}

func SetActiveNetwork(network Network) {
	activeNetwork = network
}

func GetActiveNetwork() Network {
	return activeNetwork
}