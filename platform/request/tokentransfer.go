package request

import (
	"github.com/fearlesshyena/djtwebserver/platform"
)

type GetTokenTransfers interface {
	GetTransfers(platform.ContractToken) (TokenTransfer, error)
	platform.TokenExists
}

type TokenTransfer struct {
	Transfers []platform.Transfer `json:"transfers"`
}

func (t *Token) GetTransfers(contoken platform.ContractToken) (TokenTransfer, error) {
	if transfers, err := t.Repo.GetTransfers(contoken); err != nil {
		return TokenTransfer{}, err
	} else {
		return TokenTransfer{
			Transfers: transfers,
		}, nil
	}
}
