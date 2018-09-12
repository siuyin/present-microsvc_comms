package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	//010 OMIT
	pb "github.com/siuyin/present-microsvc_comms/grpc/arith"
)

const address = "localhost:50051"

//go:generate protoc -I ../arith --go_out=plugins=grpc:../arith ../arith/arith.proto
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewArithClient(conn) // HL

	// Contact the server and print out its response.
	reply, err := c.Sum(context.Background(), &pb.SumArgs{A: 3, B: 4}) // HL
	if err != nil {
		log.Fatalf("could not compute sum: %v", err)
	}
	fmt.Printf("the sum of 3 and 4 is: %v\n", reply.Reply)
}

//020 OMIT
