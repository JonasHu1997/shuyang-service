package api

// GetHandlerFuncMap 获取处理函数
func (r *NormalAPI) GetHandlerFuncMap() map[string]func() {
	return r.handlers
}
