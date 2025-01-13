package main

import (
	"fmt"

	"admin-panel/internal/server"
)

func main() {

	newServer := server.NewServer()

	fmt.Printf("Starting Server")
	err := newServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start newServer: %s", err))
	}

}
