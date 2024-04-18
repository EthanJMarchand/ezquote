package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethanjmarchand/ezquote/internal/api"
	"github.com/ethanjmarchand/ezquote/internal/startup"
	"github.com/ethanjmarchand/ezquote/internal/web"
)

func main() {
	config, err := startup.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = run(config)
	if err != nil {
		log.Fatal(err)
	}

}

func run(s *startup.Config) error {
	service := api.NewAPIService(api.GeckoWorker{
		Config: s,
	})

	http.Handle("GET /", &web.HomeHandler{Service: *service})
	http.Handle("GET /quote/{quote}", &web.ListHandler{Service: *service})
	fmt.Printf("Starting server on port %s", s.Server)
	err := http.ListenAndServe(s.Server, nil)
	if err != nil {
		return err
	}
	return nil
}
