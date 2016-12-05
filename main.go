package main

import "github.com/surgemq/surgemq/service"

func main() {
	svr := &service.Server{}

	svr.ListenAndServe("tcp://:1883")
}
