package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("REQ RECEIVED")
		w.WriteHeader((http.StatusOK))
		jsonBytes, _ := json.Marshal(map[string]interface{}{
			"pong": true,
		})
		w.Write((jsonBytes))
	})

	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}

	fmt.Println("Server is running on localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Server not running", err)
	}
}
