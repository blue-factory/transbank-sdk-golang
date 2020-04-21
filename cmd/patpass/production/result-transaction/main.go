package main

import (
	"log"
	"os"

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
	if len(os.Args) < 2 {
		log.Fatalln("missing token arg")
	}
	token := string(os.Args[1])

	service, err := webpay.NewPatpass(privateCert, publicCert, commerceCode, commerceEmail, webpayService, webpayEnvironment)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := service.GetTransactionResult(token)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("AccountingDate", result.AccountingDate)
	log.Println("BuyOrder", result.BuyOrder)
	log.Println("CardDetail.CardNumber", result.CardDetail.CardNumber)
	log.Println("DetailOutput.SharesNumber", result.DetailOutput.SharesNumber)
	log.Println("DetailOutput.Amount", result.DetailOutput.Amount)
	log.Println("DetailOutput.CommerceCode", result.DetailOutput.CommerceCode)
	log.Println("DetailOutput.BuyOrder", result.DetailOutput.BuyOrder)
	log.Println("DetailOutput.AuthorizationCode", result.DetailOutput.AuthorizationCode)
	log.Println("DetailOutput.PaymentTypeCode", result.DetailOutput.PaymentTypeCode)
	log.Println("DetailOutput.ResponseCode", result.DetailOutput.ResponseCode)
	log.Println("SessionID", result.SessionID)
	log.Println("TransactionDate", result.TransactionDate)
	log.Println("URLRedirection", result.URLRedirection)
	log.Println("VCI", result.VCI)
}
