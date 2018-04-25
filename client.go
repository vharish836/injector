package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	gorilla "github.com/vharish836/rpc/json"
)

type getblockchaininforesp struct {
	Chain                string  `json:"chain"`
	Chainname            string  `json:"chainname"`
	Description          string  `json:"description"`
	Protocol             string  `json:"protocol"`
	Setupblocks          float64 `json:"setupblocks"`
	Blocks               float64 `json:"blocks"`
	Headers              float64 `json:"headers"`
	Bestblockhash        string  `json:"bestblockhash"`
	Difficulty           float64 `json:"difficulty"`
	Verificationprogress float64 `json:"verificationprogress"`
	Chainwork            string  `json:"chainwork"`
}

func main() {
	url := "http://localhost:6296/"
	jbuf, err := gorilla.EncodeClientRequest("getblockchaininfo", nil)
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
	var result getblockchaininforesp
	err = gorilla.DecodeClientResponse(rsp.Body, &result)
	if err != nil {
		log.Fatalf("could not decode response. %s", err)
	}
	jbuf, err = json.MarshalIndent(result, "", "\t")
	if err != nil {
		log.Fatalf("could not encode request. %s", err)
	}
	log.Printf("buffer: %s", jbuf)
	jbuf, err = gorilla.EncodeClientRequest("getblockhash", 0)
	if err != nil {
		log.Fatalf("could not encode. %s", err)
	}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jbuf))
	if err != nil {
		log.Fatalf("Could not create new request. %s", err)
	}
	req.SetBasicAuth("multichainrpc", "4FBBr9xFmiKe6RBA1F8G2ji6RBaxkQhvoQkefDx2GCK4")
	req.Header.Set("Content-Type", "application/json")
	rsp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("sending request failed. %s", err)
	}
	defer rsp.Body.Close()
	var blockhash string
	err = gorilla.DecodeClientResponse(rsp.Body, &blockhash)
	if err != nil {
		log.Fatalf("could not decode. %s", err)
	}
	log.Print(blockhash)
}
