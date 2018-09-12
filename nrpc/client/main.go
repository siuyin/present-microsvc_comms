package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/siuyin/present-microsvc_comms/nrpc"
)

//010 OMIT
func main() {
	fmt.Println("net/rpc example")

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("unable to connect to server: %v", err)
	}

	args := &nrpc.Args{3, 4}
	var reply int
	err = client.Call("Arith.Sum", args, &reply) // HL
	if err != nil {
		log.Fatalf("error encountered with rpc sum call: %v", err)
	}
	fmt.Printf("The sum of 3 and 4 is: %v\n", reply)
}

//020 OMIT
