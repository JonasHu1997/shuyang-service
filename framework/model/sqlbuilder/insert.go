package sqlbuilder

import (
	"fmt"
	"gin-service/framework/util"
	"strings"
)

func NewInsert() *SQLInsert {
	ins := new(SQLInsert)
	ins.child = ins
	return ins
}

func (ins *SQLInsert) Into(table string) ISQL {
	ins.table = table
	return ins
}

func (ins *SQLInsert) Data(data IData) ISQL {
	ins.field = data.Field()
	ins.value = data.Value()
	return ins
}

func (ins *SQLInsert) OnDuplicate(onduplicate string) ISQL {
	ins.onDuplicate = onduplicate
	return ins
}

func (ins *SQLInsert) Sql() string {
	ins.checkTable()
	tmpArr := make([]string, 0, 5)
	for _, v := range ins.field {
		tmpArr = append(tmpArr, "`"+v+"`")
	}
	field := strings.Join(tmpArr, ",")
	tmpArr = make([]string, 0, 5)
	for _, v := range ins.value {
		for i := range v {
			v[i] = "'" + util.AddSlashes(v[i]) + "'"
		}
		tmpArr = append(tmpArr, "("+strings.Join(v, ",")+")")
	}
	values := strings.Join(tmpArr, ",")
	sql := fmt.Sprintf("insert into %s (%s) values %s", ins.table, field, values)
	if ins.onDuplicate != "" {
		sql += fmt.Sprintf(" ON DUPLICATE KEY UPDATE %s", ins.onDuplicate)
	}
	return sql
}
