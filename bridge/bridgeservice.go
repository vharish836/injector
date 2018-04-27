package bridge

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"

	gorilla "github.com/vharish836/rpc/json"
)

type BridgeService struct {
	Username string
	Password string
}

func (b *BridgeService) CheckAuth(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if ok != true || username != b.Username || password != b.Password {
		return errors.New("Unauthorized")
	}
	return nil
}

func (b *BridgeService) GetResponseFromPlatform(method string, args interface{}, resp interface{}) error {
	url := "http://localhost:6296/"
	var jbuf []byte
	var err error
	pv := reflect.ValueOf(args).Elem()
	if pv.Kind() == reflect.Struct {
		if numFields := pv.NumField(); numFields == 0 {
			jbuf, err = gorilla.EncodeClientRequest(method, nil)
		} else {
			var params []interface{}
			for index := 0; index < numFields; index++ {
				params = append(params, pv.Field(index).Addr().Interface())
			}
			jbuf, err = gorilla.EncodeClientRequest(method, params)
		}
	} else {
		params := []interface{}{args}
		jbuf, err = gorilla.EncodeClientRequest(method, params)
	}
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

func (b *BridgeService) HandleAPI(r *http.Request, args interface{}, reply interface{}, method string) error {
	err := b.CheckAuth(r)
	if err != nil {
		return err
	}
	err = b.GetResponseFromPlatform(method, args, reply)
	if err != nil {
		return err
	}
	return nil
}

func (b *BridgeService) GetInfo(r *http.Request, args *GetInfoReq, reply *GetInfoRsp) error {
	return b.HandleAPI(r, args, reply, "getinfo")
}

func (b *BridgeService) GetBlockChainParams(r *http.Request, args *GetBlockChainParamsReq, reply *GetBlockChainParamsRsp) error {
	return b.HandleAPI(r, args, reply, "getblockchainparams")
}

func (b *BridgeService) GetRuntimeParams(r *http.Request, args *GetRuntimeParamsReq, reply *GetRuntimeParamsRsp) error {
	return b.HandleAPI(r, args, reply, "getruntimeparams")
}

func (b *BridgeService) SetRuntimeParam(r *http.Request, args *SetRuntimeParamReq, reply *SetRuntimeParamRsp) error {
	return b.HandleAPI(r, args, reply, "setruntimeparam")
}

func (b *BridgeService) Help(r *http.Request, args *HelpReq, reply *HelpRsp) error {
	return b.HandleAPI(r, args, reply, "help")
}

func (b *BridgeService) Stop(r *http.Request, args *StopReq, reply *StopRsp) error {
	return b.HandleAPI(r, args, reply, "stop")
}

func (b *BridgeService) AddMultiSigAddress(r *http.Request, args *AddMultiSigAddressReq, reply *AddMultiSigAddressRsp) error {
	return b.HandleAPI(r, args, reply, "addmultisigaddress")
}

func (b *BridgeService) GetAddresses(r *http.Request, args *GetAddressesReq, reply *GetAddressesRsp) error {
	return b.HandleAPI(r, args, reply, "getaddresses")
}

func (b *BridgeService) GetNewAddress(r *http.Request, args *GetNewAddressReq, reply *GetNewAddressRsp) error {
	return b.HandleAPI(r, args, reply, "getnewaddress")
}

func (b *BridgeService) ImportAddress(r *http.Request, args *ImportAddressReq, reply *ImportAddressRsp) error {
	return b.HandleAPI(r, args, reply, "importaddress")
}

func (b *BridgeService) ListAddresses(r *http.Request, args *ListAddressesReq, reply *ListAddressesRsp) error {
	return b.HandleAPI(r, args, reply, "listaddresses")
}

func (b *BridgeService) CreateKeyPairs(r *http.Request, args *CreateKeyPairsReq, reply *CreateKeyPairsRsp) error {
	return b.HandleAPI(r, args, reply, "createkeypairs")
}

func (b *BridgeService) CreateMultiSig(r *http.Request, args *CreateMultiSigReq, reply *CreateMultiSigRsp) error {
	return b.HandleAPI(r, args, reply, "createmultisig")
}

func (b *BridgeService) ValidateAddress(r *http.Request, args *ValidateAddressReq, reply *ValidateAddressRsp) error {
	return b.HandleAPI(r, args, reply, "validateaddress")
}