package def

import (
	"database/sql"
)

type ParserType uint8

type SqlType uint8

const (
	SqlTypeSelect SqlType = iota
	SqlTypeInsert
	SqlTypeUpdate
	SqlTypeDelete
)

const (
	ParserTypeQuery ParserType = iota
	ParserTypeExec
)

type SQLResult struct {
	Rows   *sql.Rows
	Result sql.Result
}

type ISQLExecor interface {
	SqlExec(sql string) (*SQLResult, error)
	SqlQuery(sql string) (*SQLResult, error)
}

type ISQLParser interface {
	New(ParserType, *SQLResult) ISQLParser
	// Bind(interface{})
	Map() map[string]string
	MapMulti() []map[string]string
}
