package transactions

import (
	"github.com/icon-project/goloop/server/jsonrpc"
	"github.com/icon-project/goloop/server/v3"
	// "paulrouge/go-icon-sdk/networks"
	// "math/big"
	// "paulrouge/go-icon-sdk/util"
)

func CallBuilder(to string, method string, params interface{}) *v3.CallParam {
	// convert to to jsonrpc.Address
	toAddress := jsonrpc.Address(to)

	// build data
	data := map[string]interface{}{
		"method": method,
	}

	if params != nil {
		data["params"] = params
	}

	callParams := v3.CallParam{
		FromAddress: "hx0000000000000000000000000000000000000000",
		ToAddress: toAddress,
		DataType: "call",
		Data: data,
		// Value: util.BigIntToHex(amount),
	}


	// txParams := v3.TransactionParam{
	// 	FromAddress: "hx9c13cd371aed69c79870b3a3f7492c10122f0315",
	// 	ToAddress: toAddress,
	// 	StepLimit: "0xf4240",
	// 	NetworkID: networks.GetActiveNetwork().NID,
	// 	Nonce: "0x1",
	// 	Version: "0x3",
	// 	Timestamp: "0x",
	// 	DataType: "call",
	// 	Data: data,
	// }

	return &callParams
}
	


