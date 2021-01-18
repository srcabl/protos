#!/bin/bash

find ./ -name "*.proto" -print | xargs -n 1 protoc --go_out=. --go_opt=module=github.com/srcabl/protos