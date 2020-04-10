package conf

import "encoding/json"

// JSONConf JSON类型的配置
type JSONConf json.RawMessage

// JSONRoot JSON类型配置 根节点
type JSONRoot map[string]json.RawMessage

// MySQL 支持的MYSQL配置
type MySQL struct {
	User         string
	Passwd       string
	Host         string
	Port         string
	Dbname       string
	Charset      string
	MaxOpenConns int
}

// IMySQL 支持转化为mysql配置要实现的接口
type IMySQL interface {
	BindMySQL(*MySQL) error
}
