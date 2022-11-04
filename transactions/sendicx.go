package transactions

import (
	"github.com/icon-project/goloop/server/jsonrpc"
	"github.com/icon-project/goloop/server/v3"
	"paulrouge/go-icon-sdk/networks"	
)

// amount is number of icx as a string
func TransferICXBuilder(to string, amount string) *v3.TransactionParam {

	// convert to to jsonrpc.Address
	toAddress := jsonrpc.Address(to)
	
	// convert amount to jsonrpc.HexInt
	amountHex := jsonrpc.HexInt(amount)

	txParams := v3.TransactionParam{
		FromAddress: "hx9c13cd371aed69c79870b3a3f7492c10122f0315",
		ToAddress: toAddress,
		Value: amountHex,
		StepLimit: "0xf4240",
		NetworkID: networks.GetActiveNetwork().NID,
		Nonce: "0x1",
		Version: "0x3",
		Timestamp: "0x",
	}

	return &txParams
}