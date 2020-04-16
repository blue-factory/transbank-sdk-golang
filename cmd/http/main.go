package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	webpayHTTP "github.com/microapis/transbank-sdk-golang/pkg/http"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

func main() {
	// read port environment value
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env must be defined")
	}

	// initialize router
	router := mux.NewRouter()

	// initialize webpay plus normal
	plusSrv, err := webpay.NewPlusNormal(privateCert, publicCert, commerceCode, commerceEmail, webpayService, webpayEnvironment)
	if err != nil {
		log.Fatalln(err)
	}

	webpayHTTP.Routes(router, &webpayHTTP.Services{
		PlusNormal: plusSrv,
		Patpass:    webpay.NewIntegrationPatpass(),
	})

	// logger middleware
	router.Use(loggingMiddleware)

	// enable cors support
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// start http server
	log.Println("Starting HTTP service...")
	go func() {
		log.Println(fmt.Sprintf("HTTP service running, Listening on port=%v", port))
		err = http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
