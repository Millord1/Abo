package entities

import "github.com/edgedb/edgedb-go"

type MintLog struct {
	id       edgedb.OptionalUUID
	Date     string   `edgedb:"date" form:""`
	TxHash   string   `edgedb:"txhash" form:""`
	contract Contract `edgedb:"contract" form:""`
	asset    Asset    `edgedb:"asset" form:""`
}
