package etherium

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {}

func Connect(address string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(address)
	if err != nil {
		return nil, err
	} else {
		return client, nil
	}
}
