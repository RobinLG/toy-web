package v1

import "sync"

// 用于确保HandlerBasedOnMap肯定实现了这个接口
var _ Handler = &HandlerBasedOnMap{}

type HandlerBasedOnMap struct {
	handlers sync.Map
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {

}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handlerFunc handlerFunc) {

}

func NewHandlerBasedOnMap() *HandlerBasedOnMap {
	return &HandlerBasedOnMap{}
}