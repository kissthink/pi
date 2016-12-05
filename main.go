package main

import (
	"github.com/surgemq/surgemq/service"
	"github.com/surgemq/surgemq/auth"
	"errors"
)

const login = "foo"
const pass = "bar"

type MyAuth struct {}

func (m MyAuth) Authenticate(id string, cred interface{}) error {

	if id == login && cred.(string) == pass {
		return nil
	}

	return errors.New("Wrong cred")
}

func main() {
	m := MyAuth{}
	provider := auth.Authenticator(m)
	auth.Register("my", provider)

	svr := &service.Server{
		Authenticator: "my",
	}

	svr.ListenAndServe("tcp://:1883")
}
