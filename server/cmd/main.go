package main

import (
	restserver "indexed-mail-search/server/pkg/rest_server"
)

func main() {
	restserver.NewRestServer().RunServer()
}
