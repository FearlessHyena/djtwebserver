package request

import (
	"github.com/fearlesshyena/djtwebserver/platform"
	"log"
)

type GetTokenOwner interface {
	GetOwner(platform.ContractToken) (TokenOwner, error)
	platform.TokenExists
}

type TokenOwner struct {
	Owner string `json:"owner"`
}

type Token struct {
	platform.Repo
}

func (t *Token) GetOwner(contoken platform.ContractToken) (TokenOwner, error) {
	if owner, err := t.Repo.GetOwner(contoken); err != nil {
		log.Fatalln(err)
		return TokenOwner{}, err
	} else {
		return TokenOwner{
			Owner: owner,
		}, nil
	}
}

func (t *Token) TokenExists(contoken platform.ContractToken) bool {
	return t.Repo.TokenExists(contoken)
}
