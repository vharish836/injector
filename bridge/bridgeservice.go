package bridge

import (
	"bytes"	
	"errors"	
	"net/http"
	"reflect"

	gorilla "github.com/vharish836/rpc/json"
)

// Service ...
type Service struct {
	Username string
	Password string
}

// CheckAuth ...
func (b *Service) CheckAuth(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if ok != true || username != b.Username || password != b.Password {
		return errors.New("Unauthorized")
	}
	return nil
}

// GetResponseFromPlatform ...
func (b *Service) GetResponseFromPlatform(method string, args interface{}, resp interface{}) error {
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

// GetInfo ...
func (b *Service) GetInfo(r *http.Request, args *GetInfoReq, reply *GetInfoRsp) error {
	return b.HandleAPI(r, args, reply, "getinfo")
}

// GetBlockChainParams ...
func (b *Service) GetBlockChainParams(r *http.Request, args *GetBlockChainParamsReq, reply *GetBlockChainParamsRsp) error {
	return b.HandleAPI(r, args, reply, "getblockchainparams")
}

// GetRuntimeParams ...
func (b *Service) GetRuntimeParams(r *http.Request, args *GetRuntimeParamsReq, reply *GetRuntimeParamsRsp) error {
	return b.HandleAPI(r, args, reply, "getruntimeparams")
}

// SetRuntimeParam ...
func (b *Service) SetRuntimeParam(r *http.Request, args *SetRuntimeParamReq, reply *SetRuntimeParamRsp) error {
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

// AddMultiSigAddress ...
func (b *Service) AddMultiSigAddress(r *http.Request, args *AddMultiSigAddressReq, reply *AddMultiSigAddressRsp) error {
	return b.HandleAPI(r, args, reply, "addmultisigaddress")
}

// GetAddresses ...
func (b *Service) GetAddresses(r *http.Request, args *GetAddressesReq, reply *GetAddressesRsp) error {
	return b.HandleAPI(r, args, reply, "getaddresses")
}

// GetNewAddress ...
func (b *Service) GetNewAddress(r *http.Request, args *GetNewAddressReq, reply *GetNewAddressRsp) error {
	return b.HandleAPI(r, args, reply, "getnewaddress")
}

// ImportAddress ...
func (b *Service) ImportAddress(r *http.Request, args *ImportAddressReq, reply *ImportAddressRsp) error {
	return b.HandleAPI(r, args, reply, "importaddress")
}

// ListAddresses ...
func (b *Service) ListAddresses(r *http.Request, args *ListAddressesReq, reply *ListAddressesRsp) error {
	return b.HandleAPI(r, args, reply, "listaddresses")
}

// CreateKeyPairs ...
func (b *Service) CreateKeyPairs(r *http.Request, args *CreateKeyPairsReq, reply *CreateKeyPairsRsp) error {
	return b.HandleAPI(r, args, reply, "createkeypairs")
}

// CreateMultiSig ...
func (b *Service) CreateMultiSig(r *http.Request, args *CreateMultiSigReq, reply *CreateMultiSigRsp) error {
	return b.HandleAPI(r, args, reply, "createmultisig")
}

// ValidateAddress ...
func (b *Service) ValidateAddress(r *http.Request, args *ValidateAddressReq, reply *ValidateAddressRsp) error {
	return b.HandleAPI(r, args, reply, "validateaddress")
}
