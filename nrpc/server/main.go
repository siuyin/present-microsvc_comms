package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"siuyin/present-microsvc_comms/nrpc"
)

//010 OMIT
func main() {
	fmt.Println("net/rpc Arith server")
	arith := new(nrpc.Arith) // nrpc from import "siuyin/present-microsvc_comms/nrpc" // HL
	rpc.Register(arith)
	rpc.HandleHTTP()
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("unable to listen on port 1234: %v", err)
	}

	go http.Serve(lis, nil)
	select {} // wait forever
}

//020 OMIT
