package platform

import "math/big"

type TokenExists interface {
	TokenExists(contoken ContractToken) bool
}

type Repo interface {
	GetOwner(contoken ContractToken) (owner string, err error)
	GetTransfers(contoken ContractToken) (transfers []Transfer, err error)
	TokenExists
}

type Transfer struct {
	Block uint64
	From string
	To string
}

type ContractToken struct {
	Contract string  `uri:"contract" binding:"required"`
	TokenId  big.Int `uri:"tokenid" binding:"required"`
}

