package router

import (
	"gin-service/framework/api"

	"github.com/gin-gonic/gin"
)

// RegRestAPI 注册RestAPI路由
func (gr *Group) RegRestAPI(fac api.IRestAPIFactory) {
	g := (*gin.RouterGroup)(gr)
	g.GET(":par", func(ctx *gin.Context) {
		api := fac.Create()
		api.SetCtx(ctx)
		api.Get()
	})
	g.GET("", func(ctx *gin.Context) {
		api := fac.Create()
		api.SetCtx(ctx)
		api.Get()
	})
	g.POST("", func(ctx *gin.Context) {
		api := fac.Create()
		api.SetCtx(ctx)
		api.Post()
	})
	g.PUT(":par", func(ctx *gin.Context) {
		api := fac.Create()
		api.SetCtx(ctx)
		api.Put()
	})
	g.DELETE(":par", func(ctx *gin.Context) {
		api := fac.Create()
		api.SetCtx(ctx)
		api.Delete()
	})
}

// RegRestAPIFunc 通过工厂函数注册RestAPI路由
func (gr *Group) RegRestAPIFunc(f func() api.IRestAPI) {
	gr.RegRestAPI(api.RestAPIFactoryFunc(f))
}

// RegNormalAPI 通过工厂类注册普通API路由
func (gr *Group) RegNormalAPI(fac api.INormalAPIFactory) {
	g := (*gin.RouterGroup)(gr)
	g.Any("/:action", func(ctx *gin.Context) {
		api := fac.Create()
		handlers := api.GetHandlerFuncMap()
		if fun := handlers[ctx.Param("action")]; fun != nil {
			api.SetCtx(ctx)
			fun()
			return
		}
		ctx.String(404, "404")
	})
}

// RegNormalAPIFunc 通过工厂函数注册普通API路由
func (gr *Group) RegNormalAPIFunc(f func() api.INormalAPI) {
	gr.RegNormalAPI(api.NormalAPIFactoryFunc(f))
}
