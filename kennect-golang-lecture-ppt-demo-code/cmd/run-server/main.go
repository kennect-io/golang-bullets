package main

import (
	"kennect-golang-lecture/customserver"
	"log"
)

func main() {
	h1 := "127.0.0.1:5000"
	log.Println("Listening on", h1)
	customserver.NewServer(h1).Listen()
}
