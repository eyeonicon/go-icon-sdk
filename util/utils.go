package util

import (
	"math/big"
	"github.com/icon-project/goloop/server/jsonrpc"
	"github.com/shopspring/decimal"
	"github.com/ubiq/go-ubiq/common/hexutil"

)

// ICX to 18 decimal loop
func ICXToLoop(iamount interface{}) *big.Int {
	decimals := 18
	amount := decimal.NewFromFloat(0)
	
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	loop := new(big.Int)
	loop.SetString(result.String(), 10)

	return loop
}

func BigIntToHex(bi *big.Int) jsonrpc.HexInt {
	hex := hexutil.EncodeBig(bi)
	_hex := jsonrpc.HexInt(hex)
	return _hex
}