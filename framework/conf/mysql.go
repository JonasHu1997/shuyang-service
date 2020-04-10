package conf

// MySQLDefault 返回默认mysql配置
func MySQLDefault() *MySQL {
	return &MySQL{
		Host:         "127.0.0.1",
		Port:         "3306",
		Charset:      "utf8",
		MaxOpenConns: 1000,
	}
}
