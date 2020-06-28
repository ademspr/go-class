package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ademspr/go-class/api/handler"
	"github.com/ademspr/go-class/config"
	"github.com/gorilla/mux"
)

type Server struct {
	config config.ServerConfiguration
	router *mux.Router
}

func (s *Server) Start(c config.ServerConfiguration) {
	s.config = c

	r := mux.NewRouter().StrictSlash(true)
	s.router = r
	s.registerRouters()

	s.run()

}

func (s *Server) run() {

	log.Printf("Listen on port:%d\n", s.config.Port)

	http_server := &http.Server{
		Addr:           fmt.Sprintf(":%d", s.config.Port),
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println(http_server.ListenAndServe())
}

func (s *Server) registerRouters() {
	s.get("/", handler.HomeHandler)
	s.get("/health", handler.HealthHandler)
}

func (s *Server) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.router.HandleFunc(path, f).Methods("GET")
}
