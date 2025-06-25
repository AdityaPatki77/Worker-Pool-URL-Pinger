package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const numWorkers = 3

var wg sync.WaitGroup

// jobs    <-chan string   // Read-only channel --> recieve only channel
// results chan<- string   // Write-only channel --> send only channel

func worker(id int, jobs <-chan string, results chan<- string) {

	defer wg.Done()

	for url := range jobs {

		fmt.Printf("Worker %d started checking %s\n", id, url)

		status := checkURL(url)

		results <- fmt.Sprintf("Worker %d -> %s -> %s", id, url, status)

		time.Sleep(2 * time.Second)
	}
}

// send a HEAD request to check if URL is alive
func checkURL(url string) string {

	resp, err := http.Head(url)

	if err != nil {
		return "ERROR: " + err.Error()
	}

	return resp.Status
}

func main() {

	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://httpstat.us/503",
		"https://example.com",
		"https://invalid.go.dev",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/200",
	}

	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))

	// launch fixed no. of workers/goroutines
	for i := 1; i <= numWorkers; i++ {

		wg.Add(1)

		go worker(i, jobs, results)
	}

	// send urls to jobs channel
	for _, url := range urls {

		jobs <- url
	}
	close(jobs)

	/// wait for workers to finish
	wg.Wait()
	close(results)

	// printing results
	fmt.Println("\nResults:")
	for res := range results {
		fmt.Println(res)
	}
}
