package framework

import (
	"net/http"
)

// Core 框架核心
type Core struct {
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}
