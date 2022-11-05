# go-icon-sdk

The Icon SDK for Go is a library for building applications on the ICON network.

## Create Client
In src/main.go in the main function:

1. Set the node you want to connect to globally. You can add networks in the networks/networks.go file.
```go
networks.SetActiveNetwork(networks.Lisbon())
```

2. Create client
```go
Client := client.NewClientV3(networks.GetActiveNetwork().URL)
```

[Click here to see all the available methods on the created Client](https://pkg.go.dev/github.com/icon-project/goloop@v1.2.14/client#NewClientV3)


## Create Wallet
When creating a new wallet it is automatically __saved as a keystore file.__ Call the function below with the _"path + filename"_. The password is used to encrypt the keystore file.

```go
wallet.CreateNewWalletAndKeystore("mywallets/keystore01", "password")
```

## Load Wallet
When loading a wallet you need to provide the path to the keystore file and the password to decrypt the keystore file.

```go
Wallet := wallet.LoadWalletFromKeystore("mywallets/keystore01", "password")
```
__Note:__ To prevent confusing between the created wallet instance and the wallet-package we name the wallet that we load "Wallet" (so with a capital W, instead of the package name).

## Send ICX
Use the TransferICXBuilder to get a transaction object. The address should be a string and the amount must be converted to a big.Int before sending it to the builder. We do this by using the "util.ICXToLoop()" function.


```go
address := "hx0000000000000000000000000000000000000000"
amount := 1 // can also by string "1" or float 1.0
bn := util.ICXToLoop(amount)

txobject := transactions.TransferICXBuilder(address, bn)
```



