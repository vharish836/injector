package bridge

import (
	"github.com/golang/glog"
	"github.com/gorilla/rpc"
	rjson "github.com/gorilla/rpc/json"
)

// Init ...
func Init(c *Config) *rpc.Server {
	s := rpc.NewServer()
	s.RegisterCodec(rjson.NewCodec(), "application/json")
	bridge := &Service{UserName: c.UserName,
		PassWord: c.PassWord,
		RPCUser: c.MultiChain.RPCUser,
		RPCPassWord: c.MultiChain.RPCPassword,
		RPCPort: c.MultiChain.RPCPort}
	err := s.RegisterService(bridge, "")
	if err != nil {
		glog.Fatalf("could not register. %s", err)
	}
	if c.EasyAPI == true {
		glog.V(1).Infof("Interceting requests for easy API support")
		s.RegisterInterceptFunc(bridge.Interceptor)
	} else {
		glog.V(1).Infof("No easy API support, use full names\n")
	}
	
	return s
}
