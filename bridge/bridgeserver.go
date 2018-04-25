package main

import (
	"log"
	"errors"
	"bytes"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	gorilla "github.com/vharish836/rpc/json"
	"net/http"
	"encoding/json"
)

type Getinforeq struct {}

type Getinforsp struct {
		Version string `json:"version"`
		Nodeversion int `json:"nodeversion"`
		Protocolversion int `json:"protocolversion"`
		Chainname string `json:"chainname"`
		Description string `json:"description"`
		Protocol string `json:"protocol"`
		Port int `json:"port"`
		Setupblocks int `json:"setupblocks"`
		Nodeaddress string `json:"nodeaddress"`
		Burnaddress string `json:"burnaddress"`
		Incomingpaused bool `json:"incomingpaused"`
		Miningpaused bool `json:"miningpaused"`
		Walletversion int `json:"walletversion"`
		Balance float64 `json:"balance"`
		Walletdbversion int `json:"walletdbversion"`
		Reindex bool `json:"reindex"`
		Blocks int `json:"blocks"`
		Timeoffset int `json:"timeoffset"`
		Connections int `json:"connections"`
		Proxy string `json:"proxy"`
		Difficulty float64 `json:"difficulty"`
		Testnet bool `json:"testnet"`
		Keypoololdest int `json:"keypoololdest"`
		Keypoolsize int `json:"keypoolsize"`
		Paytxfee float64 `json:"paytxfee"`
		Relayfee float64 `json:"relayfee"`
		Errors string `json:"errors"`
}

type Bridgeservice struct{}

func (b *Bridgeservice) Info(r *http.Request, args *Getinforeq, reply *Getinforsp) error {
	username, password, ok := r.BasicAuth()
	if ok != true || username != "username" || password != "password" {
		return errors.New("Unauthorized")
	}
	url := "http://localhost:6296/"
	jbuf, err := gorilla.EncodeClientRequest("getinfo", nil)
	if err != nil {
		log.Fatalf("could not encode. %s", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jbuf))
	if err != nil {
		log.Fatalf("Could not create new request. %s", err)
	}
	req.SetBasicAuth("multichainrpc", "4FBBr9xFmiKe6RBA1F8G2ji6RBaxkQhvoQkefDx2GCK4")
	req.Header.Set("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("sending request failed. %s", err)
	}
	defer rsp.Body.Close()
	var result Getinforsp
	err = gorilla.DecodeClientResponse(rsp.Body, &result)
	if err != nil {
		log.Fatalf("could not decode response. %s", err)
	}
	jbuf, err = json.MarshalIndent(result, "", "\t")
	if err != nil {
		log.Fatalf("could not encode request. %s", err)
	}
	log.Printf("buffer: %s", jbuf)	
	*reply = result
	return nil
}

func main()  {
	s := rpc.NewServer()
	s.RegisterCodec(gorilla.NewCodec(), "application/json")
	bridge := new(Bridgeservice)
	err := s.RegisterService(bridge,"")
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

