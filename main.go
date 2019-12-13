package main

import (
	"fmt"

	"github.com/ryanyogan/go-blog/dbclient"
	"github.com/ryanyogan/go-blog/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDB()
	service.DBClient.Seed()
}
