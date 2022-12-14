package wallet

import (
	"github.com/icon-project/goloop/module"
	"io/ioutil"
)
// Public function that creates and saves a new wallet as a keystore file to filepath.
func CreateNewWalletAndKeystore(filepath string, password string) {
	Wallet := new()
	pw := []byte(password)	
	keyStoreFromWallet(Wallet, pw, filepath)
}

// Public function that loads a wallet from a keystore file.
func LoadWallet(filepath string, password string) module.Wallet {
	pw := []byte(password)
	
	// read keystore from filepath
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	ks, err := newFromKeyStore(data,pw)
	if err != nil {
		panic(err)
	}
	
	return ks
	
}