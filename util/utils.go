// Package util provides some utility functions for the go-icon-sdk.

package util

import (
	"github.com/icon-project/goloop/server/jsonrpc"
	"github.com/shopspring/decimal"
	"github.com/ubiq/go-ubiq/common/hexutil"
	"math/big"
)

// ICX to 18 decimal Loop (Loop is the smalles unit of ICX, like Wei is the smallest unit of ETH)
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

// convert a big.Int to jsonrpc.HexInt
func BigIntToHex(bi *big.Int) jsonrpc.HexInt {
	hex := hexutil.EncodeBig(bi)
	_hex := jsonrpc.HexInt(hex)
	return _hex
}

// convert a jsonrpc.HexInt to big.Int
func HexToBigInt(hex string) *big.Int {
	bi := hexutil.MustDecodeBig(string(hex))
	return bi
}
