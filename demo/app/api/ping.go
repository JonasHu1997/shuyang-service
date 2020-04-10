package api

import (
	"gin-service/framework/api"
)

type Ping struct {
	api.RestAPI
}

func NewPing() api.IRestAPI {
	return new(Ping)
}

func (ping *Ping) Get() {
	ping.Response(200, nil, "pong")
}
