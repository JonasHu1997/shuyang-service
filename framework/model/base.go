package model

import (
	"database/sql"
	"gin-service/framework/db"
	"gin-service/framework/model/sqlbuilder"
	"gin-service/framework/model/sqlexecor"
	"gin-service/framework/model/sqlparser"
)

type Base struct{}
type SQLInitor struct{ DB *sql.DB }

func (b *Base) DB() *SQLInitor {
	return &SQLInitor{DB: db.DB}
}

func (r *SQLInitor) Select(field ...string) sqlbuilder.ISQL {
	b := sqlbuilder.NewSelect(field...)
	b.SetEP(sqlexecor.New(r.DB), new(sqlparser.SQLParser))
	return b
}

func (r *SQLInitor) Insert() sqlbuilder.ISQL {
	b := sqlbuilder.NewInsert()
	b.SetEP(sqlexecor.New(r.DB), new(sqlparser.SQLParser))
	return b
}

func (r *SQLInitor) Update(table string) sqlbuilder.ISQL {
	b := sqlbuilder.NewUpdate(table)
	b.SetEP(sqlexecor.New(r.DB), new(sqlparser.SQLParser))
	return b
}

func (r *SQLInitor) Delete() sqlbuilder.ISQL {
	b := sqlbuilder.NewDelete()
	b.SetEP(sqlexecor.New(r.DB), new(sqlparser.SQLParser))
	return b
}
