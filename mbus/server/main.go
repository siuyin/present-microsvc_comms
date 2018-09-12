package main

import (
	"fmt"
	"log"

	"github.com/nats-io/go-nats"
	"github.com/siuyin/present-microsvc_comms/mbus"
)

//010 OMIT
func main() {
	fmt.Println("message bus server")
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("unable to connect to nats server: %v", err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("unable to create encoded connection: %v", err)
	}
	defer c.Close()

	//030 OMIT
	c.Subscribe(mbus.ArithSum, func(subj, reply string, args *mbus.Args) { // HL
		c.Publish(reply, mbus.Reply(args.A+args.B)) // HL
	})
	//040 OMIT

	select {} // wait forever

}

//020 OMIT
