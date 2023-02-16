package main

// Documentation: https://go.dev/doc/diagnostics#profiling

// To generate a CPU and memory profile:
// execute 'go run main.go -cpuprofile=cpu.prof -memprofile=mem.prof'

// To show CPU and memory profile:
// install the Graphviz visualization tool. https://graphviz.org/download/
// execute 'go tool pprof -http=:8080 cpu.prof' to show CPU profile
// execute 'go tool pprof -http=:8080 mem.prof' to show memory profile

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"

	"indexed-mail-search/server/pkg/datasource"
	"indexed-mail-search/server/pkg/service"
)

func IndexEmails() {
	httpClient := &http.Client{}
	datasourceZincSearch := datasource.NewZincsearchClient(httpClient)
	indexEmailService := service.NewIndexerService(datasourceZincSearch)

	emailUsers, err := indexEmailService.GetMailUsers()

	if err != nil {
		return
	}

	var wg sync.WaitGroup
	for _, emailUser := range emailUsers {
		wg.Add(1)
		go indexEmailByUser(emailUser, &wg)
	}
	wg.Wait()
}

func indexEmailByUser(emailUser string, wg *sync.WaitGroup) {
	httpClient := &http.Client{}
	datasourceZincSearch := datasource.NewZincsearchClient(httpClient)
	indexEmailService := service.NewIndexerService(datasourceZincSearch)

	defer wg.Done()
	emails, err := indexEmailService.ProcessMailsByUser(emailUser)
	if err != nil {
		return
	}
	err = indexEmailService.IndexEmails(emails)
	if err != nil {
		return
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	IndexEmails()
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
