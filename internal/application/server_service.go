package application

import (
	"net/http"

	server "github.com/nanoHB/fcloud/internal/domain/server"
)

func InitServer() *http.Server {
	return server.NewServer()
}
