package injector

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	gorilla "github.com/vharish836/rpc/json"
)

type Bridgeservice struct {
	Username string
	Password string
}

func (b *Bridgeservice) CheckAuth(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if ok != true || username != b.Username || password != b.Password {
		return errors.New("Unauthorized")
	}
	return nil
}

func (b *Bridgeservice) SendRPCRequest(method string, args interface{}, resp interface{}) error {
	url := "http://localhost:6296/"
	jbuf, err := gorilla.EncodeClientRequest(method, args)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jbuf))
	if err != nil {
		return err
	}
	req.SetBasicAuth("multichainrpc", "4FBBr9xFmiKe6RBA1F8G2ji6RBaxkQhvoQkefDx2GCK4")
	req.Header.Set("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	err = gorilla.DecodeClientResponse(rsp.Body, &resp)
	if err != nil {
		return err
	}
	jbuf, err = json.MarshalIndent(resp, "", "\t")
	if err != nil {
		return err
	}
	log.Printf("buffer: %s", jbuf)
	return nil
}

func (b *Bridgeservice) Info(r *http.Request, args *Getinforeq, reply *Getinforsp) error {
	err := b.CheckAuth(r)
	if err != nil {
		return err
	}
	var result Getinforsp
	err = b.SendRPCRequest("getinfo", nil, &result)
	if err != nil {
		return err
	}
	*reply = result
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(gorilla.NewCodec(), "application/json")
	bridge := &Bridgeservice{Username: "username", Password: "password"}
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
