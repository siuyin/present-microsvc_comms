package main

import (
	"fmt"
	"log"
	"time"

	"siuyin/present-microsvc_comms/mbus"

	"github.com/nats-io/go-nats"
)

//010 OMIT
func main() {
	fmt.Println("message bus client")
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("unable to connect to nats server: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("unable to create encoded connection: %v", err)
	}
	defer c.Close()

	var reply mbus.Reply
	err = c.Request(mbus.ArithSum, mbus.Args{3, 4}, &reply, 1*time.Second)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	fmt.Printf("the sum of 3 and 4 is: %v", reply)
}

//020 OMIT
