package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/microapis/transbank-sdk-golang/pkg/configuration"
	"github.com/microapis/transbank-sdk-golang/pkg/service"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

func main() {
	amount := float64(1000)
	sessionID := "mi-id-de-sesion"
	buyOrder := strconv.Itoa(rand.Intn(99999))
	returnURL := "https://callback/resultado/de/transaccion"
	finalURL := "https://callback/final/post/comprobante/webpay"

	c := configuration.GetIntegrationPlusNormal()
	w := webpay.New(c)
	s := service.GetPlusNormal(w)

	transaction, err := s.InitTransaction(service.ParamsPlusNormal{
		Amount:    amount,
		SessionID: sessionID,
		BuyOrder:  buyOrder,
		ReturnURL: returnURL,
		FinalURL:  finalURL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL", transaction.URL)
	log.Println("Token", transaction.Token)
}
