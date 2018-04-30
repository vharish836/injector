package main

import (
	"flag"
	"net/http"

	"github.com/vharish836/injector/bridge"

	"github.com/gorilla/mux"
	"github.com/golang/glog"
)

func main() {
	file := flag.String("config","injector.toml","Config file")
	flag.Parse()
	c,err := bridge.GetConfig(*file)
	if err != nil {
		glog.Fatalf("Could not load config: %s\n", err)
	}
	s := bridge.Init(c)
	r := mux.NewRouter()
	r.Handle("/", s)
	err = http.ListenAndServe("localhost:8383", r)
	if err != nil {
		glog.Fatalf("could not listen. %s", err)
	}
}
