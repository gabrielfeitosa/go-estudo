package main

import (
	"fmt"
	"go-examples/core-demo/httpclient"
	"net"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://localhost:8080/timeout", nil)
	_, err := httpclient.Send(req)

	if err, ok := err.(net.Error); ok && err.Timeout() {
		fmt.Println("timeout", err)
	} else {
		fmt.Println("não timeout", err)
	}
}
