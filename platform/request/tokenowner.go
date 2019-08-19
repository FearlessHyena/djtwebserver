package request

import (
	"github.com/fearlesshyena/djtwebserver/platform"
	"log"
	"math/big"
)

type ContractReq struct {
	Contract string  `uri:"contract" binding:"required"`
	TokenId  big.Int `uri:"tokenid" binding:"required"`
}

type GetOwner interface {
	Get(ContractReq) (Token, error)
}

type Token struct {
	Owner string `json:"owner"`
}

type TokenOwner struct {
	platform.Repo
}

func (t *TokenOwner) Get(conreq ContractReq) (Token, error) {
	owner, err := t.Repo.Get(conreq.Contract, conreq.TokenId)
	if err != nil {
		log.Fatalln(err)
		return Token{}, err
	} else {
		return Token{
			Owner: owner,
		}, nil
	}
}
