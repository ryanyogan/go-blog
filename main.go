package main

import (
	"fmt"

	"github.com/ryanyogan/go-blog/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
