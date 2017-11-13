# Microservice communications

## Note on vendoring grpc
1. You *can* vendor gRPC but it will conflict with the gRPC already on your GOPATH.
1. On a clean go install the vendored gRPC will be used as it is the only one present.
