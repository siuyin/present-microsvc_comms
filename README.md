# Microservice communications

## Note on vendoring grpc
1. You *can* vendor gRPC but it will conflict with the gRPC already on your GOPATH.
1. On a clean go install the vendored gRPC will be used as it is the only one present.

## If running in a docker container
1. docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
1. setup .bashrc to have ~/go/bin in PATH
1. present -http 0.0.0.0:3999 -orighost 192.168.99.100
