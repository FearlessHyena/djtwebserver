package cache

import (
	"github.com/fearlesshyena/djtwebserver/platform"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const Selectqry = "SELECT block,from_addr, to_addr FROM token_transfer WHERE contract_addr=$1 AND token_id=$2"
const Insertqry = `INSERT INTO token_transfer (contract_addr, token_id, block, from_addr, to_addr) VALUES ($1, $2, $3, $4, $5)`

type PgClient struct {
	db *sqlx.DB
}

func NewPgClient(datasource string) (*PgClient, error) {
	if pgdb, err := sqlx.Connect("postgres", datasource);
		err != nil {
		return &PgClient{}, err
	} else {
		return &PgClient{db: pgdb}, nil
	}
}

func (c *PgClient) GetTransfers(contoken platform.ContractToken) (transfers []platform.CacheTokenTransfer, err error) {
	err = c.db.Select(&transfers,
		Selectqry,
		contoken.Contract, contoken.TokenId.Uint64())
	return
}

func (c *PgClient) WriteTransfers(contoken platform.ContractToken, tokentrans []platform.CacheTokenTransfer) (err error) {
	for i := range tokentrans {
		_, err = c.db.Exec(
			Insertqry,
			contoken.Contract, contoken.TokenId.Uint64(), tokentrans[i].Block, tokentrans[i].From, tokentrans[i].To)
	}
	return
}
