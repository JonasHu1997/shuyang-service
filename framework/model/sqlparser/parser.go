package sqlparser

import (
	"gin-service/framework/model/def"
	"gin-service/framework/util/reflect"
)

type SQLParser struct {
	t      def.ParserType
	result *def.SQLResult
}

func (s *SQLParser) New(p def.ParserType, result *def.SQLResult) def.ISQLParser {
	return &SQLParser{p, result}
}

func (s *SQLParser) MapMulti() []map[string]string {
	if s.t == def.ParserTypeExec {
		panic("only support query")
	}
	rows := s.result.Rows
	columns, _ := rows.Columns()
	data := make([]map[string]string, 0, 5)
	fCount := len(columns)
	fieldPtr := make([]interface{}, fCount)
	fieldVal := make([]string, fCount)
	for k := range columns {
		s := new(string)
		fieldPtr[k] = s
		fieldVal[k] = *s
	}

	for s.result.Rows.Next() {
		_ = rows.Scan(fieldPtr...)
		m := make(map[string]string, fCount)
		for i, v := range fieldVal {
			m[columns[i]] = v
		}
		data = append(data, m)
	}
	return data
}

func (s *SQLParser) Map() map[string]string {
	if s.t == def.ParserTypeExec {
		panic("only support query")
	}
	return s.MapMulti()[0]
}

func (s *SQLParser) Bind(dat interface{}) {
	if s.t != def.ParserTypeQuery {
		panic("only support query")
	}
	refStru := reflect.NewRefStruct(dat, "db")
	tagPtrs := refStru.TagPtr()
	rows := s.result.Rows
	columns, _ := rows.Columns()
	fieldPtr := make([]interface{}, len(columns))
	for i, v := range columns {
		fieldPtr[i] = tagPtrs[v]
	}
	s.result.Rows.Next()
	_ = rows.Scan(fieldPtr...)
}
