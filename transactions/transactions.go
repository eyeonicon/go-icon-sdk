package transactions

import (
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
	


