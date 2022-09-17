package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

func FirePingRequest(url string, client *http.Client) {
	req, _ := http.NewRequest("GET", url, nil)
	_, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func FirePingRequestWithTrace(url string, client *http.Client) {

	clientTrace := &httptrace.ClientTrace{
		GotConn: func(gci httptrace.GotConnInfo) {
			fmt.Printf("connection reused ? %t\n", gci.Reused)
		},
	}

	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	req, _ := http.NewRequestWithContext(traceCtx, "GET", url, nil)
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer response.Body.Close()
	io.Copy(io.Discard, response.Body) // if you do not wants to consume reponse received from server
	// bA, _ := io.ReadAll(response.Body) //if you wants to consume reponse received from server
	// fmt.Println("RESPONSE RECEIVED: ", string(bA))

}

var wg sync.WaitGroup

func main() {
	start := time.Now()

	hc := http.Client{
		Timeout: 10 * time.Second,
	}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		// fmt.Println(i)
		go func() {
			FirePingRequestWithTrace("http://127.0.0.1:3000/ping", &hc)
			wg.Done()
		}()
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("1000 request took %s", elapsed)
}
