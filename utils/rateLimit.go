package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var url string
var requests int
var threads int

func init() {
	log.SetFlags(0)
	log.SetPrefix("INFO: ")

	if len(os.Args) != 4 {
		log.Fatalln("Provide the right arguments in order: website-url, requests, threads")
	}

	url = os.Args[1]
	var err error
	requests, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	threads, err = strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalln(err)
	}
}

func worker(requestsChan <-chan int, wg *sync.WaitGroup) {
	for id := range requestsChan {
		startTime := time.Now()
		res, err := http.Get(url)
		if err != nil {
			log.Println("REQ", id, err)
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("REQ", id, err)
			continue
		}
		log.Println("REQ", id, time.Since(startTime).Round(time.Second), strings.TrimSpace(string(body)))
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	startTime := time.Now()
	requestsChan := make(chan int, requests)

	for range threads {
		wg.Add(1)
		go worker(requestsChan, &wg)
	}
	for i := range requests {
		requestsChan <- i
	}
	close(requestsChan)

	wg.Wait()

	duration := int(time.Since(startTime).Round(time.Second).Seconds())
	if duration > 0 {
		requestsPerSecond := int(requests / duration)
		log.Printf("Load test completed in %ds with %d rps\n", duration, requestsPerSecond)
	}
}
