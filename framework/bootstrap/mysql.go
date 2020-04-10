package bootstrap

import (
	"gin-service/framework/conf"
	"gin-service/framework/db"
)

func InitMySQL(file, path string) (err error) {
	root, err := conf.ReadJSONConf(file, path)
	if err != nil {
		return
	}
	mySQLCnf, err := root.GetConf("mysql")
	if err != nil {
		return
	}
	err = db.InitMySQL(mySQLCnf)
	if err != nil {
		return
	}
	return
}
