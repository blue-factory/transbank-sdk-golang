package main

import (
	"log"

	"github.com/microapis/transbank-sdk-golang/pkg/configuration"
	"github.com/microapis/transbank-sdk-golang/pkg/service"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

func main() {
	token := "ee0316ccc90711dfccb793547733f7be7cd4a3ac1f9e5a0e6a86a021edefaf5f"

	c := configuration.GetIntegrationPlusNormal()
	w := webpay.New(c)
	s := service.GetPlusNormal(w)

	result, err := s.GetTransactionResult(token)
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
