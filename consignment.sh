#! /bin/bash

protoc --go_out=. --go_opt=paths=source_relative \
    --micro_out=$GOPATH/src \
 ./proto/consignment/consignment.proto

#  protoc --go_out=. --go_opt=paths=source_relative \
#     --go-grpc_out=. --go-grpc_opt=paths=source_relative \   
#  ./proto/consignment/consignment.proto