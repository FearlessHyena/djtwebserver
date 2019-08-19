package main

import (
	"./token"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatalln(err)
	}

	address := common.HexToAddress("0xcc62564D40C06e2Be1F84287b0d8F6B734c856D30xcc62564D40C06e2Be1F84287b0d8F6B734c856D3")
	instance, err := token.NewDjtHashToken(address, client)
	if err != nil {
		log.Fatalln(err)
	}
	owner, err := instance.OwnerOf(nil, big.NewInt(1))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Token owner is", owner.Hex())
}
