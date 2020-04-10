package conf

import (
	"encoding/json"
	"fmt"
	"gin-service/framework/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadJSONConf 从JSON文件中读取配置
// 返回配置json的根节点
func ReadJSONConf(filename string, absPath string) (rst JSONRoot, err error) {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	// 默认从应用目录/conf下读取
	// 有传abspath 则从abs下读取
	confPath := ""
	if absPath != "" {
		confPath = absPath
	} else {
		confPath = filepath.Join(appPath, "conf")
	}
	if !util.FileExists(filepath.Join(confPath, filename)) {
		return nil, fmt.Errorf("%s Not found", filepath.Join(confPath, filename))
	}
	confData, err := ioutil.ReadFile(filepath.Join(confPath, filename))
	if err != nil {
		return
	}
	rst = make(JSONRoot)
	if err = json.Unmarshal(confData, &rst); err != nil {
		rst = nil
		return
	}
	return
}

// GetConf 从读取到的json文件中 获取值
func (root JSONRoot) GetConf(name string) (cnf JSONConf, err error) {
	if conf, ok := root[name]; ok == true {
		return JSONConf(conf), nil
	}
	return nil, fmt.Errorf("Config of %s not found", name)
}

// BindMySQL 把配置转换为mysql配置结构体
func (cnf JSONConf) BindMySQL(mysql *MySQL) error {
	return json.Unmarshal([]byte(cnf), mysql)
}
