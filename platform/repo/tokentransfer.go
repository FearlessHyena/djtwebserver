package repo

import (
	"github.com/fearlesshyena/djtwebserver/platform"
	"log"
)

type GetTokenTransfers interface {
	GetTransfers(platform.ContractToken) (TokenTransfer, error)
	platform.TokenExists
}

type TokenTransfer struct {
	Transfers []platform.Transfer `json:"transfers"`
}

func (t *Token) GetTransfers(contoken platform.ContractToken) (TokenTransfer, error) {
	//Check the cache for the record first and if it doesn't exist get it from the Repo
	cachedtrans, err := t.Cache.GetTransfers(contoken)
	if err != nil {
		log.Println("Error reading from cache. Will try to read from the repo directly", err)
	}

	if len(cachedtrans) > 0 {
		return TokenTransfer{
			Transfers: convertFromCache(cachedtrans),
		}, nil
	} else {
		log.Println("Couldn't find any records in the cache. Will try to read from the repo directly")
		if transfers, err := t.Repo.GetTransfers(contoken); err != nil {
			return TokenTransfer{}, err
		} else if err := t.Cache.WriteTransfers(contoken, convertToCache(transfers)); err != nil {
			return TokenTransfer{}, err
		} else {
			return TokenTransfer{
				Transfers: transfers,
			}, nil
		}
	}
}

func convertFromCache(cachedtrans []platform.CacheTokenTransfer) (tokentrans []platform.Transfer) {
	for i := range cachedtrans {
		tokentrans = append(tokentrans, platform.Transfer{
			Block: cachedtrans[i].Block,
			From:  cachedtrans[i].From,
			To:    cachedtrans[i].To,
		})
	}
	return
}

func convertToCache(tokentrans []platform.Transfer) (cachedtrans []platform.CacheTokenTransfer) {
	for i := range tokentrans {
		cachedtrans = append(cachedtrans, platform.CacheTokenTransfer{
			Block: tokentrans[i].Block,
			From:  tokentrans[i].From,
			To:    tokentrans[i].To,
		})
	}
	return
}
