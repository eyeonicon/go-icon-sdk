# go-icon-sdk

The Icon SDK for Go is a library for building applications on the ICON network.

## 1 Create Client
In src/main.go in the main function:

1. Set the node you want to connect to globally. You can add networks in the networks/networks.go file.
```go
networks.SetActiveNetwork(networks.Lisbon())
```


2. Create client
```go
Client := client.NewClientV3(networks.GetActiveNetwork().URL)
```

## Create Wallet
When creating a new wallet it is automatically saved as a keystore file. Call the function below with the path + filename you want to save the keystore file to. The password is used to encrypt the keystore file.

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
Use the TransferICXBuilder to get a transaction object.
```go


