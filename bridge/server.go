package bridge

import (
	"log"

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
		log.Fatalf("could not register. %s", err)
	}
	if c.EasyAPI == true {
		log.Printf("Interceting requests for easy API support\n")
		s.RegisterInterceptFunc(bridge.Interceptor)
	} else {
		log.Printf("No easy API support, use full names\n")
	}
	
	return s
}
