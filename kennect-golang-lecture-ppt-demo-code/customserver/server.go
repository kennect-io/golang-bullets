package customserver

import (
	"encoding/json"
	"log"
	"net/http"
)

// Server struct for storing the HTTP server instance
type Server struct {
	httpServer *http.Server
}

// NewServer returns an instance of the struct Server
func NewServer(hostAddr string) *Server {
	mux := http.NewServeMux()

	// demo ping handler
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		jsonBytes, err := json.Marshal(map[string]bool{
			"pong": true,
		})
		if err != nil {
			log.Println("error occured while json marshalling", err)
		}
		w.Write(jsonBytes)
	})

	return &Server{
		httpServer: &http.Server{
			Addr:    hostAddr,
			Handler: mux,
		},
	}
}

// Listen will start a http server with a /ping handler
func (s *Server) Listen() {

	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalln("error starting server:", err)
	}
}
