package sqlbuilder

import (
	"fmt"
	"gin-service/framework/util"
	"strings"
)

func (m MapC) ToCondStr() string {
	strArr := make([]string, 3)
	for k, v := range m {
		str := fmt.Sprintf("`%s`='%s'", k, util.AddSlashes(v))
		strArr = append(strArr, str)
	}
	return strings.Join(strArr, ",")
}

func (m C) ToCondStr() string {
	return string(m)
}
