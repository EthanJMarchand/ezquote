package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ethanjmarchand/ezquote/internal/quoting"
	"github.com/ethanjmarchand/ezquote/internal/startup"
	"github.com/ethanjmarchand/ezquote/internal/web"
)

func main() {
	config, err := startup.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	if err = run(config); err != nil {
		log.Fatal(err)
	}
}

func run(s *startup.Config) error {
	// TODO: Return an Error
	service := quoting.NewAPIService(quoting.GeckoWorker{
		Config: s,
		NetClient: &http.Client{
			Timeout: time.Second * 10,
		},
	})
	// TODO: Create "new" functions for homehandler and listhandler
	http.Handle("GET /", &web.HomeHandler{Service: *service})
	http.Handle("GET /quote/{quote}", &web.ListHandler{Service: *service})
	fmt.Printf("Starting server on port %s", s.Server)
	err := http.ListenAndServe(s.Server, nil)
	if err != nil {
		return fmt.Errorf("server closed %w", err)
	}
	return nil
}
