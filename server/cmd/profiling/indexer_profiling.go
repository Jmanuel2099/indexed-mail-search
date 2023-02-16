package indexerprofiling

// Documentation: https://go.dev/doc/diagnostics#profiling

// To generate a CPU and memory profile:
// execute 'go run main.go -cpuprofile=cpu.prof -memprofile=mem.prof'

// To show CPU and memory profile:
// install the Graphviz visualization tool. https://graphviz.org/download/
// execute 'go tool pprof -http=:8080 cpu.prof' to show CPU profile
// execute 'go tool pprof -http=:8080 mem.prof' to show memory profile

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"indexed-mail-search/server/pkg/datasource"
	"indexed-mail-search/server/pkg/service"
)

type IndexerProfiling struct {
	indexedService *service.IndexerEmailService
}

func NewIndexerProfiling() *IndexerProfiling {
	indexerProfiling := &IndexerProfiling{}
	indexerProfiling.configureServiceDependency()

	return indexerProfiling
}

func (ip *IndexerProfiling) configureServiceDependency() {
	httpClient := &http.Client{}
	datasourceZincSearch := datasource.NewZincsearchClient(httpClient)

	ip.indexedService = service.NewIndexerService(datasourceZincSearch)
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func (ip *IndexerProfiling) StartProfiling() {
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

	start := time.Now()
	ip.indexEmails()
	end := time.Now()
	fmt.Println("execution time" + end.Sub(start).String())

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

func (ip *IndexerProfiling) indexEmails() {
	emailUsers, err := ip.indexedService.GetMailUsers()
	if err != nil {
		fmt.Println("error occurred getting users in the database. " + err.Error())
		return
	}

	var wg sync.WaitGroup
	for _, emailUser := range emailUsers {
		wg.Add(1)
		go ip.indexEmailByUser(emailUser, &wg)
	}
	wg.Wait()
}

func (ip *IndexerProfiling) indexEmailByUser(emailUser string, wg *sync.WaitGroup) {
	defer wg.Done()
	emails, err := ip.indexedService.ProcessMailsByUser(emailUser)
	if err != nil {
		fmt.Println("error occurred when trying to  emails from user" + emailUser + ". " + err.Error())
		return
	}
	err = ip.indexedService.IndexEmails(emails)
	if err != nil {
		fmt.Println("error occurred when trying to index mails in zincsearch. " + err.Error())
		return
	}
}
