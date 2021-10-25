package v1

import "net/http"

type Context struct {
	w http.ResponseWriter
	Routable *http.Request
}


