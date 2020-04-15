package main

import (
	"log"

	"github.com/microapis/transbank-sdk-golang/pkg/service"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

func main() {
	token := "ea67dd4a90f2abc7d577c35dfcca5d3bd9688f528515a9f35c7408b42e01a7a"

	w := webpay.NewIntegrationPlusNormal()

	result, err := service.GetPlusNormal(w).GetTransactionResult(token)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("AccountingDate", result.AccountingDate)
	log.Println("BuyOrder", result.BuyOrder)
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
