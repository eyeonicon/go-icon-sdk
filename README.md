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





## Send ICX
Use the TransferICXBuilder to get a transaction object.
```go


