package sqlexecor

import (
	"database/sql"
	"gin-service/framework/model/def"
)

type SQLExecor struct {
	db *sql.DB
}

func New(db *sql.DB) *SQLExecor {
	s := new(SQLExecor)
	s.db = db
	return s
}

func (s *SQLExecor) SqlQuery(query string) (*def.SQLResult, error) {
	rows, err := s.db.Query(query)
	// TryCatchQuery(err, query)
	return &def.SQLResult{Rows: rows}, err
}

func (s *SQLExecor) SqlExec(exec string) (*def.SQLResult, error) {
	result, err := s.db.Exec(exec)
	return &def.SQLResult{Result: result}, err
}
