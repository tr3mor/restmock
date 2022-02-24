package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restmock/internal/app/restmock"
	"time"
)

func main() {
	l := restmock.NewLogger()
	var confPath, port, host string
	flag.StringVar(&confPath, "conf-path", "/etc/restmock/conf.yaml", "Path to yaml config with described interactions")
	flag.StringVar(&host, "host", "0.0.0.0", "Host on which server should run")
	flag.StringVar(&port, "port", "8080", "Port on which server should run")
	flag.Parse()
	conf := restmock.ParseConfig(confPath, l)
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	for _, rule := range conf.Interactions {
		r.HandleFunc(rule.Request.Path, restmock.NewHandlerFunc(rule, l)).Methods(rule.Request.Method)
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Cant start server", err)
	}
}
