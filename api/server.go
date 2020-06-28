package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartApi(port int) {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/version", VersionHandler)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
