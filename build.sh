#!/bin/bash

# ensure all previously generated files are destroyed
find ./ -name "*.go" -print | xargs -n 1 rm
# generate the files a fresh
find ./ -name "*.proto" -print | xargs -n 1 protoc --go_out=. --go-grpc_out=. --go_opt=module=github.com/srcabl/protos
#since -go-grpc_out forces everything in to the package diredctory path (or i don't know what i am doing)
#we need to manually move them and delete the directory
find ./github.com/srcabl/protos "*.go" -type f -print | while read line; do mv "${line}" ".${line#./github.com/srcabl/protos}"; done;
rm -rf ./github.com
#
# ensure go mod is all set
go mod tidy