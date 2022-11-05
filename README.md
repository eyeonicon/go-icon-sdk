# go-icon-sdk

The Icon SDK for Go is a library for building applications on the ICON network.

## Create Client
In src/main.go in the main function:

1. Set the node you want to connect to globally. You can add networks in the networks/networks.go file.
```go
// Lisbon Testnet
networks.SetActiveNetwork(networks.Lisbon())

// Mainnet
networks.SetActiveNetwork(networks.Mainnet())
```

2. Create client
```go
Client := client.NewClientV3(networks.GetActiveNetwork().URL)
```

[Click here to see all the available methods on the created Client](https://pkg.go.dev/github.com/icon-project/goloop@v1.2.14/client#NewClientV3)


## Create Wallet
When creating a new wallet it is automatically __saved as a keystore file.__ Call the function below with the _"path + filename"_. The password is used to encrypt the keystore file.

```go
wallet.CreateNewWalletAndKeystore("../mywallets/keystore01", "password")
```

## Load Wallet
When loading a wallet you need to provide the path to the keystore file and the password to decrypt the keystore file.

```go
Wallet := wallet.LoadWalletFromKeystore("../mywallets/keystore01", "password")
```
__Note:__ To prevent confusing between the created wallet instance and the wallet-package we name the wallet that we load "Wallet" (so with a capital W, instead of the package name).

## Send ICX
Use the TransferICXBuilder to get a transaction object. The address should be a string and the amount must be converted to a big.Int before sending it to the builder. We do this by using the "util.ICXToLoop()" function.


```go
// set address & amount of ICX to sent
address := "hx0000000000000000000000000000000000000000" // must be a string
amount := 1 // can also be a string "1" or a float 1.0

// convert amount of icx to loop in big.Int
bn := util.ICXToLoop(amount)

// create transaction object
txobject := transactions.TransferICXBuilder(address, bn)

// we need to have a wallet loaded to sign the transaction
Wallet := wallet.LoadWallet("../mywallets/keystore01", "password")

// sign & send the transaction
tx, err := Client.SendTransaction(Wallet, txobject)
if err != nil {
    fmt.Println(err)
}

// print the transaction hash (not working correctly atm!)
fmt.Println(tx)

```

## Call a Smart Contract on the ICON Blockchain
Use the CallBuilder to get a call-object. The Callbuilder takes in the address of the smart contract as a string, the name of the method you want to call (also as a string) and a params object. If the method you want to call does not take any parameters you can just pass in a empty object.

```go
    // set address
	a := "cx26a32e36df0a408a573163d05b3043c180359735"
	
    // set the method to call -> there is a method on this contract called "name"
    method := "name" 
	
    // this method does not take in any parameters, we do need to create an empty object
    params := map[string]interface{}{}

    // create call object
	callObject := transactions.CallBuilder(a,method, params)

    // make the call
	response, err := Client.Call(callObject)
	if err != nil {
		fmt.Println(err)
	}

    // print the response
    fmt.Println(response)
```

