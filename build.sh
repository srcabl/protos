#!/bin/bash

# ensure all previously generated files are destroyed
find ./ -name "*.go" -print | xargs -n 1 rm
# generate the files a fresh
find ./ -name "*.proto" -print | xargs -n 1 protoc --go_out=. --go_opt=module=github.com/srcabl/protos
# ensure go mod is all set
go mod tidy