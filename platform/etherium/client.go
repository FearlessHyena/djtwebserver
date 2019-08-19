package etherium

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type EthClient struct {
	Client *ethclient.Client
}

func Connect(address string) (*EthClient, error) {
	client, err := ethclient.Dial(address)
	if err != nil {
		return nil, err
	} else {
		return &EthClient{
			Client: client,
		}, nil
	}
}

func (e *EthClient) Get(conaddr string, tokenid big.Int) (string, error) {
	address := common.HexToAddress(conaddr)
	instance, err := NewDjtHashToken(address, e.Client)
	if err != nil {
		return "", err
	}
	tokenowner, err := instance.OwnerOf(nil, &tokenid)
	if err != nil {
		return "", err
	} else {
		return tokenowner.Hex(), nil
	}
}
