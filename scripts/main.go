package main

import (
	"fmt"
	"net/http"

	"github.com/go-send-message/scripts/router"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("run main")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                               // All origins
		AllowedMethods: []string{"POST, GET, OPTIONS, PUT, DELETE"}, // Allowing only get, just an example
	})

	srv := &http.Server{
		Addr:    ":8910",
		Handler: c.Handler(router.Index),
	}

	srv.ListenAndServe()
}
