package bridge

import (
	"bytes"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/vharish836/rpc"
	gorilla "github.com/vharish836/rpc/json"
)

// Service ...
type Service struct {
	UserName    string
	PassWord    string
	RPCUser     string
	RPCPassWord string
	RPCPort     string
}

// Interceptor ...
func (b *Service) Interceptor(r *rpc.RequestInfo) *http.Request {
	req := r.Request
	method := "Service." + strings.ToUpper(r.Method[:1]) + r.Method[1:]
	req.Header.Set("X-Method", method)
	return req
}

// CheckAuth ...
func (b *Service) CheckAuth(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if ok != true || username != b.UserName || password != b.PassWord {
		return errors.New("Unauthorized")
	}
	return nil
}

// GetResponseFromPlatform ...
func (b *Service) GetResponseFromPlatform(method string, args interface{}, resp interface{}) error {
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
	req, err := http.NewRequest("POST", "http://localhost:"+b.RPCPort+"/", bytes.NewBuffer(jbuf))
	if err != nil {
		return err
	}
	req.SetBasicAuth(b.RPCUser, b.RPCPassWord)
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
	return nil
}

// HandleAPI ...
func (b *Service) HandleAPI(r *http.Request, args interface{}, reply interface{}, method string) error {
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

// Getinfo ...
func (b *Service) Getinfo(r *http.Request, args *GetInfoReq, reply *GetInfoRsp) error {
	return b.HandleAPI(r, args, reply, "getinfo")
}

// Getblockchainparams ...
func (b *Service) Getblockchainparams(r *http.Request, args *GetBlockChainParamsReq, reply *GetBlockChainParamsRsp) error {
	return b.HandleAPI(r, args, reply, "getblockchainparams")
}

// Getruntimeparams ...
func (b *Service) Getruntimeparams(r *http.Request, args *GetRuntimeParamsReq, reply *GetRuntimeParamsRsp) error {
	return b.HandleAPI(r, args, reply, "getruntimeparams")
}

// Setruntimeparam ...
func (b *Service) Setruntimeparam(r *http.Request, args *SetRuntimeParamReq, reply *SetRuntimeParamRsp) error {
	return b.HandleAPI(r, args, reply, "setruntimeparam")
}

// Help ...
func (b *Service) Help(r *http.Request, args *HelpReq, reply *HelpRsp) error {
	return b.HandleAPI(r, args, reply, "help")
}

// Stop ...
func (b *Service) Stop(r *http.Request, args *StopReq, reply *StopRsp) error {
	return b.HandleAPI(r, args, reply, "stop")
}

// Addmultisigaddress ...
func (b *Service) Addmultisigaddress(r *http.Request, args *AddMultiSigAddressReq, reply *AddMultiSigAddressRsp) error {
	return b.HandleAPI(r, args, reply, "addmultisigaddress")
}

// Getaddresses ...
func (b *Service) Getaddresses(r *http.Request, args *GetAddressesReq, reply *GetAddressesRsp) error {
	return b.HandleAPI(r, args, reply, "getaddresses")
}

// Getnewaddress ...
func (b *Service) Getnewaddress(r *http.Request, args *GetNewAddressReq, reply *GetNewAddressRsp) error {
	return b.HandleAPI(r, args, reply, "getnewaddress")
}

// Importaddress ...
func (b *Service) Importaddress(r *http.Request, args *ImportAddressReq, reply *ImportAddressRsp) error {
	return b.HandleAPI(r, args, reply, "importaddress")
}

// Listaddresses ...
func (b *Service) Listaddresses(r *http.Request, args *ListAddressesReq, reply *ListAddressesRsp) error {
	return b.HandleAPI(r, args, reply, "listaddresses")
}

// Createkeypairs ...
func (b *Service) Createkeypairs(r *http.Request, args *CreateKeyPairsReq, reply *CreateKeyPairsRsp) error {
	return b.HandleAPI(r, args, reply, "createkeypairs")
}

// Createmultisig ...
func (b *Service) Createmultisig(r *http.Request, args *CreateMultiSigReq, reply *CreateMultiSigRsp) error {
	return b.HandleAPI(r, args, reply, "createmultisig")
}

// Validateaddress ...
func (b *Service) Validateaddress(r *http.Request, args *ValidateAddressReq, reply *ValidateAddressRsp) error {
	return b.HandleAPI(r, args, reply, "validateaddress")
}
