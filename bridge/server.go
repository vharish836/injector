package bridge

import (
	"log"

	"github.com/vharish836/rpc"
	gorilla "github.com/vharish836/rpc/json"
)

// Init ...
func Init(c *Config) *rpc.Server {
	s := rpc.NewServer()
	s.RegisterCodec(gorilla.NewCodec(), "application/json")
	bridge := &Service{UserName: c.UserName, 
		PassWord: c.PassWord, 
		RPCUser: c.MultiChain.RPCUser,
		RPCPassWord: c.MultiChain.RPCPassword,
		RPCPort: c.MultiChain.RPCPort}
	err := s.RegisterService(bridge, "")
	if err != nil {
		log.Fatalf("could not register. %s", err)
	}
	s.RegisterInterceptFunc(bridge.Interceptor)
	return s
}
