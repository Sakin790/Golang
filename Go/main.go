package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// fetchData fetches data from a URL and sends it to results channel
func fetchData(url string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error reading %s: %v", url, err)
		return
	}

	results <- fmt.Sprintf("Data from %s: %s", url, string(body[:50])) // first 50 chars
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/4",
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/6",
		"https://jsonplaceholder.typicode.com/todos/7",
		"https://jsonplaceholder.typicode.com/todos/8",
		"https://jsonplaceholder.typicode.com/todos/9",
		"https://jsonplaceholder.typicode.com/todos/10",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls)) // buffered channel

	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, results, &wg)
	}

	wg.Wait()
	close(results) // close channel when all goroutines finish

	// Collect all results
	finalData := []string{}
	for res := range results {
		finalData = append(finalData, res)
	}

	// Store or print final data
	for _, data := range finalData {
		fmt.Println(data)
	}
}
