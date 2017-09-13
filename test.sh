#!/bin/bash

failure=0
echo "mode: set" > coverage.out
go test -v . -covermode=count -coverprofile=coverage.out || failure=1
