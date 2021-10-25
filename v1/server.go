package v1

// Routable 可路由的
type Routable interface {
	// Route 设定一个路由器，命中该路由的会执行handlerFunc的代码
	Route(method string, pattern string, handlerFunc handlerFunc)
}

// Server 是http server的顶级抽象
type Server interface {
	Routable
	// Start 启动服务器
	Start(address string) error
}

type sdkHttpServer struct {
	// Name server的名字，可用在日志输出上
	Name string
	handler Handler
}

func (s sdkHttpServer) Route(method string, pattern string, handlerFunc handlerFunc) {
	panic("implement me")
}

func (s sdkHttpServer) Start(address string) error {
	panic("implement me")
}

func NewSdkHttpServer(name string) Server {

	handler := NewHandlerBasedOnMap()
	res := &sdkHttpServer{
		Name: name,
		handler: handler,
	}
	return res
}
