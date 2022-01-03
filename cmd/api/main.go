package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config struct holds all the configuration settings for application
type config struct {
	port int
	env  string
}

// application struct holds teh dependencies for HTTP handlers
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 9000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.port),
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
