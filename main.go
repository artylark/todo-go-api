package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
