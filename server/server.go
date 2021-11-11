package main

import (
	// "server/face"
	mnet "server/Net"
)

func main() {
	server := mnet.NewServer()
	server.Serve()

}
