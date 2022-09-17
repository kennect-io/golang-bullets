package main

import (
	"context"
	"kennect-golang-lecture/request"
	"log"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

// MakeRequestsWithNoReuseNoTrace with no TCP conn reuse no trace makes 1000 http requests
func MakeRequestsWithNoReuseNoTrace(host string, httpClient http.Client) {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		request.FirePingRequest(host, &httpClient)
	}
	elapsed := time.Since(start)
	log.Printf("1000 requests took %s", elapsed)
}

// MakeRequestWithNoReuseAndTrace with no TCP conn reuse traces the http requests makes 1000 http requests
func MakeRequestWithNoReuseAndTrace(host string, httpClient http.Client) {
	// client trace to log whether the request's underlying tcp connection was re-used
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn was reused: %t", info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	start := time.Now()
	for i := 0; i < 1000; i++ {
		request.FirePingRequestWithNoReuse(host, &httpClient, traceCtx)
	}
	elapsed := time.Since(start)
	log.Printf("1000 requests took %s", elapsed)
}

// MakeRequestWithReuseAndTrace with TCP conn reuse traces the http requests makes 1000 http requests
func MakeRequestWithReuseAndTrace(host string, httpClient http.Client) {
	// client trace to log whether the request's underlying tcp connection was re-used
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn was reused: %t", info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	start := time.Now()
	for i := 0; i < 1000; i++ {
		request.FirePingRequestWithReuse(host, &httpClient, traceCtx)
	}
	elapsed := time.Since(start)
	log.Printf("1000 requests took %s", elapsed)
}

// MakeRequestWithReuseAndTraceWithGoRoutines with TCP conn reuse traces the http requests makes 1000 http requests as 1000 go routines
func MakeRequestWithReuseAndTraceWithGoRoutines(host string, httpClient http.Client) {
	// client trace to log whether the request's underlying tcp connection was re-used
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn was reused: %t", info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			request.FirePingRequestWithReuse(host, &httpClient, traceCtx)
			wg.Done()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("1000 requests took %s", elapsed)
}

func main() {
	h1 := "127.0.0.1:5000"

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// no reuse no trace 1000 http requests
	MakeRequestsWithNoReuseNoTrace(h1, httpClient)

	// no reuse 1000 http requests
	MakeRequestWithNoReuseAndTrace(h1, httpClient)

	// reuse 1000 http requests
	MakeRequestWithReuseAndTrace(h1, httpClient)

	// reuse 1000 http requests go routines
	MakeRequestWithReuseAndTraceWithGoRoutines(h1, httpClient)

}
