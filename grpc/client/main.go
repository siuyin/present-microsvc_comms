package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	//010 OMIT
	pb "siuyin/present-microsvc_comms/grpc/arith"
)

const address = "localhost:50051"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewArithClient(conn)

	// Contact the server and print out its response.
	reply, err := c.Sum(context.Background(), &pb.SumArgs{3, 4})
	if err != nil {
		log.Fatalf("could not compute sum: %v", err)
	}
	fmt.Printf("the sum of 3 and 4 is: %v", reply.Reply)
}

//020 OMIT
