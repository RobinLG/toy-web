package main

import (
	"fmt"
	"net/http"
	v1 "toy-web/v1"
)

func main() {
	server := v1.NewSdkHttpServer("test-server")
	server.Route(http.MethodGet, "/user/signup", SignUp)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}