package sqlbuilder

import (
	"gin-service/framework/model/def"
	"strings"
)

// SetEP 设置执行器、解析器
func (base *SQLBase) SetEP(exe def.ISQLExecor, par def.ISQLParser) {
	base.exe = exe
	base.par = par
}

// From 组装select xxx from delete from
func (base *SQLBase) From(table string) ISQL {
	if base.sType != def.SqlTypeSelect && base.sType != def.SqlTypeDelete {
		panic("unsupport method From")
	}
	base.table = table
	return base.child
}

// OrderBy select子类实现
func (base *SQLBase) OrderBy(table string) ISQL {
	panic("unsupport method orderby")
}

// Limit select、update、delete的limit
func (base *SQLBase) Limit(limit int) ISQL {
	if base.sType == def.SqlTypeInsert {
		panic("unpupport method Where")
	}
	base.limit = limit
	return base.child
}

// Start select子类实现
func (base *SQLBase) Start(start int) ISQL {
	panic("unsupport method start")
}

// Data update、insert子类实现
func (base *SQLBase) Data(data IData) ISQL {
	panic("unsupport method data")
}

// Into insert子类实现
func (base *SQLBase) Into(table string) ISQL {
	panic("unsupport method From")
}

// OnDupilate insert子类实现
func (base *SQLBase) OnDuplicate(field string) ISQL {
	panic("unsupport method From")
}

// Where 组装where条件 insert语句不支持
func (base *SQLBase) Where(cond ...ICond) ISQL {
	condArr := make([]string, 5)
	for _, s := range cond {
		t := s.ToCondStr()
		if t != "" {
			condArr = append(condArr, t)
		}
	}
	if len(condArr) == 0 {
		return base.child
	}
	base.cond = strings.Join(condArr, " and ")
	return base.child
}

// Do 链式调用执行SQL 需要在初始化时设置外层定义的查询器
func (base *SQLBase) Do() def.ISQLParser {
	sql := base.child.Sql()
	if base.exe == nil {
		panic("query not set")
	}
	if base.sType == def.SqlTypeSelect {
		rst, _ := base.exe.SqlQuery(sql)
		return base.par.New(def.ParserTypeQuery, rst)
	}
	rst, _ := base.exe.SqlExec(sql)
	return base.par.New(def.ParserTypeExec, rst)
}

func (base *SQLBase) checkTable() {
	if base.table == "" {
		panic("Table Cannot Be Empty!")
	}
	base.table = "`" + base.table + "`"
}
