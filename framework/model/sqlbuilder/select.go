package sqlbuilder

import (
	"fmt"
	"strings"
)

func NewSelect(field ...string) *SQLSelect {
	sel := new(SQLSelect)
	sel.child = sel
	sel.field = field
	return sel
}

func (sel *SQLSelect) OrderBy(orderby string) ISQL {
	sel.order = orderby
	return sel
}

func (sel *SQLSelect) Start(start int) ISQL {
	sel.start = start
	return sel
}

func (sel *SQLSelect) Sql() string {
	sel.checkTable()
	field := "*"
	if len(sel.field) != 0 {
		joi := true
		for i, v := range sel.field {
			if v == "*" {
				joi = false
				break
			}
			sel.field[i] = fmt.Sprintf("`%s`", v)
		}
		if joi {
			field = strings.Join(sel.field, ",")
		}
	}

	sql := fmt.Sprintf("select %s from %s", field, sel.table)
	if sel.cond != "" {
		sql += fmt.Sprintf(" where %s", sel.cond)
	}
	if sel.limit > 0 {
		limit := ""
		if sel.start > 0 {
			limit = fmt.Sprintf(" limit %d, %d", sel.start, sel.limit)
		} else {
			limit = fmt.Sprintf(" limit %d", sel.limit)
		}
		sql += limit
	}
	return sql
}
