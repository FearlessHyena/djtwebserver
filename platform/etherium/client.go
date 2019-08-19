package etherium

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fearlesshyena/djtwebserver/platform"
	"math/big"
)

type EthClient struct {
	Client *ethclient.Client
}

func Connect(address string) (*EthClient, error) {
	if client, err := ethclient.Dial(address); err != nil {
		return nil, err
	} else {
		return &EthClient{
			Client: client,
		}, nil
	}
}

func (e *EthClient) GetOwner(contoken platform.ContractToken) (string, error) {
	if instance, err := e.getInstance(contoken); err != nil {
		return "", err
	} else if tokenowner, err := instance.OwnerOf(nil, &contoken.TokenId); err != nil {
		return "", err
	} else {
		return tokenowner.Hex(), nil
	}
}

func (e *EthClient) TokenExists(contoken platform.ContractToken) bool {
	if instance, err := e.getInstance(contoken); err != nil {
		return false
	} else if _, err := instance.TokenURI(nil, &contoken.TokenId); err != nil {
		return false
	} else {
		return true
	}
}

func (e *EthClient) GetTransfers(contoken platform.ContractToken) ([]platform.Transfer, error) {
	if instance, err := e.getInstance(contoken); err != nil {
		return nil, err
	} else {
		tokens := []*big.Int{&contoken.TokenId}
		if transiterator, err := instance.FilterTransfer(nil, nil, nil, tokens); err != nil {
			return nil, err
		} else {
			var transfers []platform.Transfer
			for ok := transiterator.Next(); ok; ok = transiterator.Next() {
				transfer := transiterator.Event
				transfers = append(transfers, platform.Transfer{
					Block: transfer.Raw.BlockNumber,
					From:  transfer.From.Hex(),
					To:    transfer.To.Hex(),
				})
			}
			return transfers, nil
		}
	}
}

func (e *EthClient) getInstance(contoken platform.ContractToken) (*DjtHashToken, error) {
	address := common.HexToAddress(contoken.Contract)
	if instance, err := NewDjtHashToken(address, e.Client); err != nil {
		return nil, nil
	}else {
		return instance, err
	}
}
