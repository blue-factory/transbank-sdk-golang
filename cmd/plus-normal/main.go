package main

import (
	"log"
	"math/rand"

	"github.com/microapis/transbank-sdk-golang/pkg/configuration"
	"github.com/microapis/transbank-sdk-golang/pkg/service"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

func main() {
	amount := float64(1000)
	sessionID := "mi-id-de-sesion"
	buyOrder := int64(rand.Intn(9999999))
	returnURL := "https://callback/resultado/de/transaccion"
	finalURL := "https://callback/final/post/comprobante/webpay"

	c := configuration.GetIntegrationPlusNormal()
	w := webpay.New(c)
	t := service.GetPlusNormal(w)

	transaction, err := t.InitTransaction(service.ParamsPlusNormal{
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
