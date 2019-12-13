#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

go mod download
go build -o accountservice-linux-amd64
echo build `pwd`

export GOOS=darwin

docker build -t ryanyogan/accountservice .

docker service rm accountservice
docker service create --name=accountservice --replicas=1 --network=blog -p=6767:6767 ryanyogan/accountservice