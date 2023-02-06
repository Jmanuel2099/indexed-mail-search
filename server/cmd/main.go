package main

import (
	"fmt"
	restserver "indexed-mail-search/server/pkg/rest_server"
)

func main() {
	fmt.Println("Application is running")

	restserver.NewRestServer().RunServer()
}
