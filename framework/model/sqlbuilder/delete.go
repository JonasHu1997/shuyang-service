package sqlbuilder

import (
	"fmt"
)

func NewDelete() *SQLDelete {
	del := new(SQLDelete)
	del.child = del
	return del
}

func (del *SQLDelete) Sql() string {
	del.checkTable()
	sql := fmt.Sprintf("delete from %s", del.table)
	if del.cond != "" {
		sql += " where " + del.cond
	}
	if del.limit > 0 {
		sql += fmt.Sprintf(" limit %d", del.limit)
	}
	return sql
}
