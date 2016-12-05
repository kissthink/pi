package mqtt

import (
	"github.com/surgemq/surgemq/auth"
	"github.com/surgemq/surgemq/service"
	"fmt"
)

var srv *service.Server

type AuthProvider struct {}

func (a AuthProvider) Authenticate(id string, cred interface{}) error {
	return nil
}

func init() {
	m := AuthProvider{}
	provider := auth.Authenticator(m)
	auth.Register("myProvider", provider)

	srv = &service.Server{
		Authenticator: "myProvider",
	}
}

func StartServer(port int) error {
	return srv.ListenAndServe(fmt.Sprintf("tcp://:%d", port))
}
