package main

import (
	"net/http"
	v1 "toy-web/v1"
)

func main() {
	server := v1.NewSdkHttpServer("test-server")
	server.Route(http.MethodGet, "/user/signup", v1.SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}

//func home(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "这是主页")
//}
//
//func user(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "这是用户")
//}
//func createUser(w http.ResponseWriter, r *http.Request) {R
//	fmt.Fprintf(w, "这是创建用户")
//}
//
//func order(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "这是订单")
//}
