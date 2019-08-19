package platform

import "math/big"

type Repo interface {
	Get(conaddr string, tokenid big.Int) (owner string, err error)
}
