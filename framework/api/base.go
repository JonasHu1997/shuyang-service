package api

// API公用方法封装
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// SetCtx 路由调用这个方法把gin的ctx嵌入
func (r *Base) SetCtx(ctx *gin.Context) {
	r.ctx = ctx
}

func (r *Base) m(code int, data interface{}, msg string) gin.H {
	resp := gin.H{
		"code":    code,
		"data":    data,
		"message": msg,
	}
	return resp
}

// Response 返回正常业务请求
// 正常业务请求定义：业务预期之内的请求（业务逻辑正常）
func (r *Base) Response(code int, data interface{}, msg string) {
	r.ctx.JSON(200, r.m(code, data, msg))
}

// ServerErr 返回服务端错误
// 服务端错误：业务预期之外的错误，如mysql插入失败等错误
func (r *Base) ServerErr(code int, data interface{}, msg string) {
	r.ctx.JSON(http.StatusInternalServerError, r.m(code, data, msg))
}

// BadRequest 返回客户端错误
// 客户端错误：参数类型错误等
func (r *Base) BadRequest(code int, data interface{}, msg string) {
	r.ctx.JSON(http.StatusBadRequest, r.m(code, data, msg))
}

// CustomResponse 自定义响应
func (r *Base) CustomResponse(httpCode, code int, data interface{}, msg string) {
	r.ctx.JSON(httpCode, r.m(code, data, msg))
}

// bodyJSON request body中的JSON转为map
func (r *Base) bodyJSON() (ret map[string]interface{}) {
	ret = make(map[string]interface{}, 0)
	err := r.BindBodyJSON(&ret)
	if err != nil {
		r.BadRequest(400, nil, "JSON parse Error")
	}
	return ret
}

// BindBodyJSON body中JSON解析成结构体
// 可以设计结构体的field为指针类型 若数据为空 bind后值为nil
func (r *Base) BindBodyJSON(i interface{}) error {
	err := r.ctx.ShouldBindJSON(i)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			r.BadRequest(400, nil, fmt.Sprintf("JSON parse Error:(%s)", err.Error()))
		default:
			r.BadRequest(400, nil, "JSON parse Error")
		}
		return err
	}
	return nil
}

// BindQueryJSON query中JSON解析成结构体
func (r *Base) BindQueryJSON(i interface{}) error {
	err := r.ctx.ShouldBindQuery(i)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			r.BadRequest(400, nil, fmt.Sprintf("JSON parse Error:(%s)", err.Error()))
		default:
			r.BadRequest(400, nil, "JSON parse Error")
		}
		return err
	}
	return nil
}

// QueryJSON  query中json解析成map
func (r *Base) QueryJSON() (ret map[string]interface{}) {
	ret = make(map[string]interface{}, 0)
	err := r.BindQueryJSON(&ret)
	if err != nil {
		r.BadRequest(400, nil, "JSON parse Error")
	}
	return ret
}
