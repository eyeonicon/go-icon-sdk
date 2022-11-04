package wallet

import (
	"github.com/icon-project/goloop/module"
	"io/ioutil"
)


func CreateNewWalletAndKeystore(filepath string, password string) {
	Wallet := New()
	pw := []byte(password)	
	KeyStoreFromWallet(Wallet, pw, filepath)
}

func LoadWallet(filepath string, password string) module.Wallet {
	pw := []byte(password)
	
	// read keystore from filepath
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	ks, err := NewFromKeyStore(data,pw)
	if err != nil {
		panic(err)
	}
	
	return ks
	
}