package sqlbuilder

import "gin-service/framework/model/def"

// ISQL 全部支持的链式操作
type ISQL interface {
	From(string) ISQL        // from xxx delete select 公用
	OrderBy(string) ISQL     // select xx from xx order by
	Where(...ICond) ISQL     // 除了insert 以外公用 逻辑相同
	Limit(int) ISQL          // 除了insert 以外公用 逻辑相同
	Start(int) ISQL          // select
	OnDuplicate(string) ISQL //insert
	Data(IData) ISQL         // update insert 公用
	Sql() string             // 所有公用
	Do() def.ISQLParser      // 所有公用
}

type SQLBase struct {
	table string
	child ISQL
	cond  string
	limit int
	sType def.SqlType
	exe   def.ISQLExecor
	par   def.ISQLParser
}

// SQLSelect SQL select语句结构体
type SQLSelect struct {
	SQLBase
	field []string
	order string
	start int
}

type SQLInsert struct {
	field       []string
	value       [][]string
	onDuplicate string
	SQLBase
}

type SQLUpdate struct {
	SQLBase
	field []string
	value []string
	data  string
}

type SQLDelete struct {
	SQLBase
}

// C where条件中直接传入字符串 用and拼接
type C string

// MapC where条件中 传入Map 转义字符 然后用and拼接
type MapC map[string]string

// ICond where条件中 允许传入的类型
type ICond interface {
	ToCondStr() string
}

// D insert update Data函数直接传入字符串
type D string

// MapD insert or update Data map
type MapD map[string]string

// MapDArr insert Data arr
type MapDArr []MapD

// MapDTransfer insert update Data函数传入map
type MapDTransfer struct {
	field []string
	value [][]string
}

type IData interface {
	Field() []string
	Value() [][]string
}
