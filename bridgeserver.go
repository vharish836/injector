package main

import (
	"strings"
	"log"
	"net/http"

	"github.com/vharish836/injector/bridge"

	"github.com/gorilla/mux"
	"github.com/vharish836/rpc"	
	gorilla "github.com/vharish836/rpc/json"
)

// Interceptor ...
func Interceptor(r *rpc.RequestInfo) *http.Request {
	req := r.Request
	method := "Service." + strings.ToUpper(r.Method[:1]) + r.Method[1:]
	req.Header.Set("X-Method",method)
	return req
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(gorilla.NewCodec(), "application/json")
	bridge := &bridge.Service{Username: "username", Password: "password"}
	err := s.RegisterService(bridge, "")
	if err != nil {
		log.Fatalf("could not register. %s", err)
	}
	s.RegisterInterceptFunc(Interceptor)	
	r := mux.NewRouter()
	r.Handle("/", s)
	err = http.ListenAndServe("localhost:8383", r)
	if err != nil {
		log.Fatalf("could not listen. %s", err)
	}
}
