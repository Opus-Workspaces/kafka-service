package modules

import (
	"fmt"
	"kafka-service/app/cmd/server"
	"kafka-service/app/routers"
)

func Modules(server *server.Server) {
	apiVersion := server.Echo.Group(routers.Ver1)
	// Using middleware if needed
	fmt.Println(apiVersion)
}
