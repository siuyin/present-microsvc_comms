package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//010 OMIT
	pb "siuyin/present-microsvc_comms/grpc/arith"
)

//go:generate protoc -I ../arith --go_out=plugins=grpc:../arith ../arith/arith.proto
const (
	port = ":50051"
)

// server is used to implement ArithServer
type server struct{}

// Sum implements ArithServer.Sum
func (s *server) Sum(ctx context.Context, in *pb.SumArgs) (*pb.SumReply, error) {
	return &pb.SumReply{in.A + in.B}, nil
}

//020 OMIT

//030 OMIT
func main() {
	fmt.Println("GRPC server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArithServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//040 OMIT
