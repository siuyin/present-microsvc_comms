# Microservice communications

## present play will not work
The play function of present compiles and runs go code
in a tree outside of the module root.
Thus play will not be able to use go.mod and thus fails to work.

## gRPC on go modules
1. This edition tries out gRPC with go modules.
  The vendor folder is ignored.

1. If you have cloned this repo in you $GOPATH/src then  
  export GO111MODULE=on to force go modules to be active.

## If running in a docker container
1. docker run -it --name present -v godata:/home/siuyin/go -p 3999:3999 siuyin/go:dev
1. setup .bashrc to have ~/go/bin in PATH or export PATH=~/go/bin:$PATH
1. present -http 0.0.0.0:3999 -orighost 192.168.99.100
