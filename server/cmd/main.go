package main

import (
	"fmt"

	indexerprofiling "indexed-mail-search/server/cmd/profiling"
	restserver "indexed-mail-search/server/pkg/rest_server"
)

func main() {
	var mode int

	fmt.Println("modes: 1 -> server. 2 -> profiling.")
	fmt.Scanln(&mode)

	if mode == 1 {
		restserver.NewRestServer().RunServer()
	} else if mode == 2 {
		indexerprofiling.NewIndexerProfiling().StartProfiling()
	} else {
		fmt.Println("Incorrect data, the options are 1 -> server. 2 -> profiling.")
	}
}
