package api

import (
	"github.com/gin-gonic/gin"
)

// Base API基类
type Base struct {
	ctx *gin.Context
}

// RestAPI 路由到此API实例 根据http method调用对应方法处理
type RestAPI struct {
	Base
}

// NormalAPI 路由到此API实例 由内部的handlers map决定路由到哪个函数
type NormalAPI struct {
	Base
	handlers map[string]func()
}

// IAPI API接口
type IAPI interface {
	SetCtx(*gin.Context)
}

// IRestAPI RestAPI接口
type IRestAPI interface {
	IAPI
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

// INormalAPI 普通API 需要有个方法调用获取内部路由
type INormalAPI interface {
	IAPI
	GetHandlerFuncMap() map[string]func()
}

// IRestAPIFactory RestAPI工厂 路由调用Create方法 返回API实例
type IRestAPIFactory interface {
	Create() IRestAPI
}

// RestAPIFactoryFunc RestAPI工厂函数 路由调用返回API实例
type RestAPIFactoryFunc func() IRestAPI

// INormalAPIFactory 普通API工厂 路由调用Create方法 返回API实例
type INormalAPIFactory interface {
	Create() INormalAPI
}

// NormalAPIFactoryFunc 普通API工厂函数 路由调用返回API实例
type NormalAPIFactoryFunc func() INormalAPI
