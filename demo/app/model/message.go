package model

import (
	"gin-service/framework/model"
	"gin-service/framework/model/sqlbuilder"
)

type Message struct {
	model.Base
}

func (m *Message) GetMessage(id int) (ret map[string]string) {
	// select * from message where id = 8 limit 1
	return m.DB().Select("*").From("message").Where(sqlbuilder.C("id = 1")).Limit(1).Do().Map()
}
