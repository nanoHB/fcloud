package main

import (
	"fmt"

	application "github.com/nanoHB/fcloud/internal/application"
)

func main() {

	server := application.InitServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
