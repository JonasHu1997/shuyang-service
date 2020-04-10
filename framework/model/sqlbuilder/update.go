package sqlbuilder

import (
	"fmt"
	"strings"
)

func NewUpdate(table string) *SQLUpdate {
	up := new(SQLUpdate)
	up.table = table
	up.child = up
	return up
}

func (up *SQLUpdate) Data(data IData) ISQL {
	up.field = data.Field()
	up.value = data.Value()[0]
	return up
}

func (up *SQLUpdate) Sql() string {
	up.checkTable()
	tmpArr := make([]string, 0, 5)
	for i := range up.field {
		tmpArr = append(tmpArr, fmt.Sprintf("set `%s`='%s'", up.field[i], up.value[i]))
	}
	sql := fmt.Sprintf("update %s %s", up.table, strings.Join(tmpArr, ","))
	if up.limit > 0 {
		sql += fmt.Sprintf(" limit %d", up.limit)
	}
	return sql
}
