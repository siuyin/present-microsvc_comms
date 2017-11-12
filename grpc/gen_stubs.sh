#!/bin/sh
# include grpc/arith folder, compile the arith.proto file, gen stubs in grpc/arith folder
protoc -I grpc/arith grpc/arith/arith.proto --go_out=plugins=grpc:grpc/arith
ls -lt grpc/arith
