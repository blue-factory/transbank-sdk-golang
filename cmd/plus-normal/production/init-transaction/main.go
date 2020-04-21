package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/microapis/transbank-sdk-golang"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

const (
	privateCert       = "private_cert_str"
	publicCert        = "public_cert_str"
	commerceCode      = 0
	commerceEmail     = "commerce_email"
	webpayService     = webpay.ServiceNormal
	webpayEnvironment = webpay.EnvironmentProduction
)

func main() {
	amount := float64(1000)
	sessionID := "mi-id-de-sesion"
	buyOrder := strconv.Itoa(rand.Intn(99999))
	returnURL := "https://callback/resultado/de/transaccion"
	finalURL := "https://callback/final/post/comprobante/webpay"

	service, err := webpay.NewPlusNormal(privateCert, publicCert, commerceCode, commerceEmail, webpayService, webpayEnvironment)
	if err != nil {
		log.Fatalln(err)
	}

	transaction, err := service.InitTransaction(transbank.InitTransaction{
		Amount:    amount,
		SessionID: sessionID,
		BuyOrder:  buyOrder,
		ReturnURL: returnURL,
		FinalURL: finalURL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL", transaction.URL)
	log.Println("Token", transaction.Token)
}
