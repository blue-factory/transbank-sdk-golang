package main

import (
	"fmt"

	webpayHTTP "github.com/microapis/transbank-sdk-golang/pkg/http"
)

func main() {
	// initialize router
	router := mux.NewRouter()

	webpayHTTP.Routes(router)

	// logger middleware
	router.Use(loggingMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Println("Starting HTTP service...")
	go func() {
		log.Println(fmt.Sprintf("HTTP service running, Listening on: %v", addr))
		err = http.ListenAndServe(":5000", handler)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
