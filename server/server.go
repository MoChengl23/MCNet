package main

import (
	"fmt"
	"server/mnet"
)

// mnet "server/Net"

func main() {

	test := test{}
	test.AddQueue()
	fmt.Println(test.GetLen())

	server := mnet.NewServer()
	server.Serve()

}
