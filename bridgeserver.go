package main

import (
	"log"
	"net/http"

	"github.com/vharish836/jrpc/injector"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	gorilla "github.com/vharish836/rpc/json"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(gorilla.NewCodec(), "application/json")
	bridge := &injector.Bridgeservice{Username: "username", Password: "password"}
	err := s.RegisterService(bridge, "")
	if err != nil {
		log.Fatalf("could not register. %s", err)
	}
	r := mux.NewRouter()
	r.Handle("/", s)
	err = http.ListenAndServe("localhost:8383", r)
	if err != nil {
		log.Fatalf("could not listen. %s", err)
	}
}
