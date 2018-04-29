package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/vharish836/injector/bridge"

	"github.com/gorilla/mux"
)

func main() {
	file := flag.String("config","injector.toml","Config file")
	flag.Parse()
	c,err := bridge.GetConfig(*file)
	if err != nil {
		log.Fatalf("Could not load config: %s\n", err)
	}
	s := bridge.Init(c)
	r := mux.NewRouter()
	r.Handle("/", s)
	err = http.ListenAndServe("localhost:8383", r)
	if err != nil {
		log.Fatalf("could not listen. %s", err)
	}
}
