package main

import (
	"fmt"
	"golang-tests-talk/internal/request"
	"golang-tests-talk/internal/request/contracts"
)

func main() {
	if do(request.NewRequest()) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failed!")
	}
}

func do(requestClient contracts.Request) bool {
	return requestClient.KeepAlive("https://www.google.com")
}
