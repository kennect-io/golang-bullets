package request

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

// FirePingRequest no trace simple request
func FirePingRequest(hostAddr string, client *http.Client) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s/ping", hostAddr), nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("request failed", err)
	}

	// connection not reused
	defer resp.Body.Close()

}

// FirePingRequestWithNoReuse tracing request and proving that connection isnt reused
func FirePingRequestWithNoReuse(hostAddr string, client *http.Client, traceContext context.Context) {
	req, _ := http.NewRequestWithContext(traceContext, "GET", fmt.Sprintf("http://%s/ping", hostAddr), nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("request failed", err)
	}

	// connection not reused
	defer resp.Body.Close()

}

// FirePingRequestWithReuse tracing and proving that connection is definitely reused
func FirePingRequestWithReuse(hostAddr string, client *http.Client, traceContext context.Context) {
	req, _ := http.NewRequestWithContext(traceContext, "GET", fmt.Sprintf("http://%s/ping", hostAddr), nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("request failed", err)
	}

	// connection will be reused
	defer resp.Body.Close()
	// io.Copy(io.Discard, resp.Body)
	io.ReadAll(resp.Body)

}
