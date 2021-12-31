package main

import "server/mnet"

// mnet "server/Net"

func main() {

	// a := new([]int)
	// a[0] = 2

	// for _, i := range a {
	// 	fmt.Println(i)
	// }
	server := mnet.NewServer()
	server.Serve()

}
