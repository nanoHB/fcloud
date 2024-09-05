package main

import (
	"fmt"

	server "github.com/nanoHB/fcloud/internal/application/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
