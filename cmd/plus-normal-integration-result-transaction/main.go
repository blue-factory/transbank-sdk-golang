package main

import (
	"log"

	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

func main() {
	token := "ec3ff8ef147754ce4ce82700cb90faebefc3841715716d364988297b019d4e47"

	service := webpay.NewIntegrationPlusNormal()
	result, err := service.GetTransactionResult(token)
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
