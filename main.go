package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func producer(total int) chan *http.Request {
	queue := make(chan *http.Request)
	// make sure that this the for-loop is run using goroutine
	// otherwise, it will block
	go func() {
		for i := 0; i < total; i++ {
			// ignore error assuming that the request valid
			req, _ := http.NewRequest("GET", "http://localhost:9000", nil)
			// enqueue the generated request
			queue <- req
		}
		// close the channel when then number of generated requests reach total
		close(queue)
	}()
	return queue
}

func consumer(queue chan *http.Request, id int) {
	// create http client
	client := &http.Client{}

	// consumer runs forever until the queue is closed
	for {
		select {
		case req, ok := <-queue:
			if !ok {
				// not ok means queue is closed
				// we stop consumer for loop here
				break
			}

			// fire the request
			resp, err := client.Do(req)
			if err != nil {
				// continue the the next request for now
				continue
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("error: %v", err)
				continue
			}

			fmt.Printf("consumer id %d, resp: %s\n", id, string(body))
		}
	}
}

func main() {
	// create fixed size queue
	queue := producer(100)

	// create consumers to consume work in the queue
	consumerCount := 10
	for i := 0; i < consumerCount; i++ {
		go consumer(queue, i)
	}

	// time.Sleep(5 * time.Second)
}
