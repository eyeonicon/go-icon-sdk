package transactions

import (
	"math/big"
	"paulrouge/go-icon-sdk/networks"
	"paulrouge/go-icon-sdk/util"
	"github.com/icon-project/goloop/module"
	"github.com/icon-project/goloop/server/jsonrpc"
	"github.com/icon-project/goloop/server/v3"
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

	return &callParams
}
	
func TransactionBuilder(from module.Address, to string, method string, params interface{}, value *big.Int) *v3.TransactionParam {
	// convert to to jsonrpc.Address
	toAddress := jsonrpc.Address(to)

	// convert from to jsonrpc.Address
	fromAddress := jsonrpc.Address(from.String())

	// build data object
	data := map[string]interface{}{
		"method": method,
	}

	if params != nil {
		data["params"] = params
	}
	
	txParams := v3.TransactionParam{
		FromAddress: fromAddress,
		ToAddress: toAddress,
		Value: util.BigIntToHex(value),
		StepLimit: "0xf4240",
		NetworkID: networks.GetActiveNetwork().NID,
		Nonce: "0x1",
		Version: "0x3",
		Timestamp: "0x",
		DataType: "call",
		Data: data,
	}

	return &txParams
}




