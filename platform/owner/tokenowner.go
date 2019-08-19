package owner

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fearlesshyena/djtwebserver/platform/etherium"
	"log"
	"math/big"
)


type ContractRequest struct {
	Contract string `uri:"contract" binding:"required"`
	TokenId big.Int `uri:"tokenid" binding:"required"`
}

type GetOwner interface {
	Get(ContractRequest) (TokenOwner, error)
}

type TokenOwner struct {
	Owner string `json:"owner"`
}

type Repo struct {
	Client *ethclient.Client
}

func New() (*Repo, error) {
	client, err := etherium.Connect("https://ropsten.infura.io")
	if err != nil {
		return nil, err
	} else {
		return &Repo{
			Client: client,
		}, nil
	}
}

func (r *Repo) Get(conreq ContractRequest) (TokenOwner, error) {
	address := common.HexToAddress(conreq.Contract)
	instance, err := etherium.NewDjtHashToken(address, r.Client)
	if err != nil {
		log.Fatalln(err)
		return TokenOwner{}, err
	}
	owner, err := instance.OwnerOf(nil, &conreq.TokenId)
	if err != nil {
		log.Fatalln(err)
		return TokenOwner{}, err
	}

	return TokenOwner{
		Owner: owner.Hex(),
	}, nil
}
