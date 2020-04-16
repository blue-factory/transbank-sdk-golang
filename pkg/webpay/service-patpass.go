package webpay

import (
	"encoding/xml"

	"github.com/microapis/transbank-sdk-golang"
)

const (
	patpassTransactionType = "TR_NORMAL_WS_WPM"
)

type patpass struct {
	webpay *Webpay
}

// NewPatpass returns a Webpay Plus Normal with respective configuration
func NewPatpass(privateCert, publicCert string, commerceCode int64, commerceEmail, service, environment string) (transbank.Transaction, error) {
	w, err := New(privateCert, publicCert, commerceCode, commerceEmail, service, environment)
	if err != nil {
		return nil, err
	}

	return &patpass{
		webpay: w,
	}, nil
}

// NewIntegrationPatpass returns a configured Webpay instance that will use
// the integration environment
func NewIntegrationPatpass() transbank.Transaction {
	return &patpass{
		webpay: new(getIntegrationPatpass()),
	}
}

// InitTransaction performans a "patpass" transaction and returns a token
func (pp *patpass) InitTransaction(params transbank.InitTransaction) (*transbank.InitTransactionResponse, error) {
	// TODO(ca): missign implementation for check if params are valid

	bodyRequest := patpassInitTransactionBodyRequest{
		ID:        "_0",
		XMLnsSOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		TnsInitTransaction: initTransactionResquest{
			XMLnsTns:          "http://service.wswebpay.webpay.transbank.com/",
			WSTransactionType: patpassTransactionType,
			CommerceCode:      pp.webpay.GetCommerceCode(),
			Amount:            params.Amount,
			SessionID:         params.SessionID,
			ReturnURL:         params.ReturnURL,
			FinalURL:          params.FinalURL,
			BuyOrder:          params.BuyOrder,
			DetailBuyOrder:    params.BuyOrder,

			WPMDetail: &patpassWPMDetailRequest{
				ServiceID:           params.WPMDetail.ServiceID,
				CardHolderID:        params.WPMDetail.CardHolderID,
				CardHolderName:      params.WPMDetail.CardHolderName,
				CardHolderLastName1: params.WPMDetail.CardHolderLastName1,
				CardHolderLastName2: params.WPMDetail.CardHolderLastName2,
				CardHolderMail:      params.WPMDetail.CardHolderMail,
				CellPhoneNumber:     params.WPMDetail.CellPhoneNumber,
				ExpirationDate:      params.WPMDetail.ExpirationDate,
				CommerceMail:        params.WPMDetail.CommerceMail,
				UfFlag:              params.WPMDetail.UfFlag,
			},
		},
	}

	return baseInitTransaction(pp.webpay, bodyRequest)
}

// GetTransactionResult validates a transaction given a token
func (pp *patpass) GetTransactionResult(token string) (*transbank.TransactionResultResponse, error) {
	return baseGetTransactionResult(pp.webpay, token)
}

type patpassInitTransactionBodyRequest struct {
	XMLName            xml.Name `xml:"soap:Body"`
	XMLnsSOAP          string   `xml:"xmlns:soap,attr,omitempty"`
	ID                 string   `xml:"Id,attr,omitempty"`
	TnsInitTransaction initTransactionResquest
}

type patpassWPMDetailRequest struct {
	XMLName xml.Name `xml:"wPMDetail"`
	ServiceID           string `xml:"serviceId"`
	CardHolderID        string `xml:"cardHolderId"`
	CardHolderName      string `xml:"cardHolderName"`
	CardHolderLastName1 string `xml:"cardHolderLastName1"`
	CardHolderLastName2 string `xml:"cardHolderLastName2"`
	CardHolderMail      string `xml:"cardHolderMail"`
	CellPhoneNumber     string `xml:"cellPhoneNumber"`
	ExpirationDate      string `xml:"expirationDate"`
	CommerceMail        string `xml:"commerceMail"`
	UfFlag              bool   `xml:"ufFlag"`
}
