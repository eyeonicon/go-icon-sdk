package wallet

import (
	"github.com/icon-project/goloop/common"
	"github.com/icon-project/goloop/common/crypto"
	"github.com/icon-project/goloop/module"
)

type softwareWallet struct {
	skey *crypto.PrivateKey
	pkey *crypto.PublicKey
}

// returns the address of the wallet
func (w *softwareWallet) Address() module.Address {
	return common.NewAccountAddressFromPublicKey(w.pkey)
}

// signs the data with the private key of the wallet
func (w *softwareWallet) Sign(data []byte) ([]byte, error) {
	sig, err := crypto.NewSignature(data, w.skey)
	if err != nil {
		return nil, err
	}
	return sig.SerializeRSV()
}

// returns the public key of the wallet
func (w *softwareWallet) PublicKey() []byte {
	return w.pkey.SerializeCompressed()
}

// creates a new key pair and returns a new wallet
func new() module.Wallet {
	sk, pk := crypto.GenerateKeyPair()
	return &softwareWallet{
		skey: sk,
		pkey: pk,
	}
}

// creates a key pair from a private key
func newFromPrivateKey(sk *crypto.PrivateKey) (module.Wallet, error) {
	pk := sk.PublicKey()
	return &softwareWallet{
		skey: sk,
		pkey: pk,
	}, nil
}
