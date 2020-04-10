package api

// Create 实现INormalAPI接口
func (f NormalAPIFactoryFunc) Create() INormalAPI {
	return f()
}

// Create 实现IRestAPI接口
func (f RestAPIFactoryFunc) Create() IRestAPI {
	return f()
}
