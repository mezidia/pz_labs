package main

import (
	"fmt"
	"net/http"

	"github.com/mezidia/pz_labs/lab_3/server/tools"
)

type HttpHandlerFunc http.HandlerFunc

type Router struct {
	routs map[string]HttpHandlerFunc
}

func (r *Router) getHandler(route string) HttpHandlerFunc {
	return r.routs[route]
}

func (r *Router) addRoute(route string, handler HttpHandlerFunc) {
	r.routs[route] = handler
}

func (r *Router) rmRoute(route string) {
	delete(r.routs, route)
}

func (r *Router) handle(rw http.ResponseWriter, req *http.Request) {
	handler, ok := r.routs[req.RequestURI]
	fmt.Println(handler)
	fmt.Println(req.RequestURI)
	if !ok {
		tools.WriteJsonBadRequest(rw, "404")
	} else {
		handler(rw, req)
	}
}
