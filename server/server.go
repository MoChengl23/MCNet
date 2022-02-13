package main

import (
	"server/mnet"

	"fmt"
)

// mnet "server/Net"

func main() {

	a := mnet.NewTestStruct()
	a.Test()
	a.Print()
	

	// server := mnet.NewServer()
	// server.Serve()

}
func DoSomething() {
	fmt.Println("I do something")
}
