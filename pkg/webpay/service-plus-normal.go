package webpay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/url"

	"github.com/microapis/transbank-sdk-golang/pkg/sign"

	"github.com/microapis/transbank-sdk-golang"
)

const (
	plusNormalTransactionType = "TR_NORMAL_WS"
)

type plusNormal struct {
	webpay *Webpay
}

// NewPlusNormal returns a Webpay Plus Normal with respective configuration
func NewPlusNormal(privateCert, publicCert string, commerceCode int64, commerceEmail, service, environment string) (transbank.Transaction, error) {
	w, err := New(privateCert, publicCert, commerceCode, commerceEmail, service, environment)
	if err != nil {
		return nil, err
	}

	return &plusNormal{
		webpay: w,
	}, nil
}

// NewIntegrationPlusNormal returns a configured Webpay instance that will use
// the integration environment
func NewIntegrationPlusNormal() transbank.Transaction {
	return &plusNormal{
		webpay: new(getIntegrationPlusNormal()),
	}
}

// InitTransaction performans a "plusNormal" transaction and returns a token
// func (pn *plusNormal) InitTransaction(amount float64, sessionID, buyOrder, returnURL, finalURL string) (*transbank.ResponsePlusNormalInitTransaction, error) {
func (pn *plusNormal) InitTransaction(params transbank.InitTransaction) (*transbank.InitTransactionResponse, error) {
	if params.Amount <= 0 {
		return nil, errors.New("invalid Amount")
	}
	if params.SessionID == "" {
		return nil, errors.New("undefined SessionID")
	}
	if params.BuyOrder == "" {
		return nil, errors.New("undefined BuyOrder")
	}
	u, err := url.Parse(params.FinalURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, errors.New("invalid FinalURL")
	}
	u, err = url.Parse(params.ReturnURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return nil, errors.New("invalid ReturnURL")
	}

	bodyRequest := initTransactionBodyRequest{
		ID:        "_0",
		XMLnsSOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		TnsInitTransaction: initTransactionResquest{
			XMLnsTns:          "http://service.wswebpay.webpay.transbank.com/",
			SessionID:         params.SessionID,
			ReturnURL:         params.ReturnURL,
			FinalURL:          params.FinalURL,
			CommerceCode:      pn.webpay.GetCommerceCode(),
			Amount:            params.Amount,
			BuyOrder:          params.BuyOrder,
			DetailBuyOrder:    params.BuyOrder,
			WSTransactionType: plusNormalTransactionType,
		},
	}

	return baseInitTransaction(pn.webpay, bodyRequest)
}

// GetTransactionResult validates a transaction given a token
func (pn *plusNormal) GetTransactionResult(token string) (*transbank.TransactionResultResponse, error) {
	return baseGetTransactionResult(pn.webpay, token)
}

func baseInitTransaction(webpay *Webpay, body interface{}) (*transbank.InitTransactionResponse, error) {
	b, err := webpay.SOAP(body)
	if err != nil {
		return nil, err
	}

	res := &initTransactionEnvolpeResponse{}
	err = xml.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	if res.Body.Fault != nil {
		errMsg := fmt.Sprintf("Error: code=%s message=%s", res.Body.Fault.Code, res.Body.Fault.Message)
		return nil, errors.New(errMsg)
	}

	it := res.Body.Ns2InitTransactionResponse

	return &transbank.InitTransactionResponse{
		URL:   it.URL,
		Token: it.Token,
	}, nil
}

func baseGetTransactionResult(webpay *Webpay, token string) (*transbank.TransactionResultResponse, error) {
	if token == "" {
		return nil, errors.New("invalid token")
	}

	bodyRequest := transactionResultBodyRequest{
		ID:        "_0",
		XMLnsSOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		TnsAcknowledgeTransaction: transactionResultResquest{
			XMLnsTns:   "http://service.wswebpay.webpay.transbank.com/",
			TokenInput: token,
		},
	}

	b, err := webpay.SOAP(bodyRequest)
	if err != nil {
		return nil, err
	}

	res := &transactionResultEnvolpeResponse{}
	err = xml.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	if res.Body.Fault != nil {
		errMsg := fmt.Sprintf("Error: code=%s message=%s", res.Body.Fault.Code, res.Body.Fault.Message)
		return nil, errors.New(errMsg)
	}

	tr := res.Body.Ns2TransactionResultResponse

	return &transbank.TransactionResultResponse{
		AccountingDate: tr.AccountingDate,
		BuyOrder:       tr.BuyOrder,
		CardDetail: transbank.CardDetail{
			CardNumber: tr.CardNumber,
		},
		DetailOutput: transbank.DetailOutput{
			SharesNumber:      tr.SharesNumber,
			Amount:            tr.Amount,
			CommerceCode:      tr.CommerceCode,
			BuyOrder:          tr.DetailBuyOrder,
			AuthorizationCode: tr.AuthorizationCode,
			PaymentTypeCode:   tr.PaymentTypeCode,
			ResponseCode:      tr.ResponseCode,
		},
		SessionID:       tr.SessionID,
		TransactionDate: tr.TransactionDate,
		URLRedirection:  tr.URLRedirection,
		VCI:             tr.VCI,
	}, nil
}

type initTransactionBodyRequest struct {
	XMLName            xml.Name `xml:"soap:Body"`
	XMLnsSOAP          string   `xml:"xmlns:soap,attr,omitempty"`
	ID                 string   `xml:"Id,attr,omitempty"`
	TnsInitTransaction initTransactionResquest
}

type initTransactionResquest struct {
	XMLName           xml.Name `xml:"tns:initTransaction"`
	XMLnsTns          string   `xml:"xmlns:tns,attr,omitempty"`
	WSTransactionType string   `xml:"wsInitTransactionInput>wSTransactionType"`
	SessionID         string   `xml:"wsInitTransactionInput>sessionId"`
	ReturnURL         string   `xml:"wsInitTransactionInput>returnURL"`
	FinalURL          string   `xml:"wsInitTransactionInput>finalURL"`
	BuyOrder          string   `xml:"wsInitTransactionInput>buyOrder"`

	CommerceCode   int64   `xml:"wsInitTransactionInput>transactionDetails>commerceCode"`
	Amount         float64 `xml:"wsInitTransactionInput>transactionDetails>amount"`
	DetailBuyOrder string  `xml:"wsInitTransactionInput>transactionDetails>buyOrder"`

	WPMDetail *patpassWPMDetailRequest `xml:"wsInitTransactionInput>wPMDetail"`
}

type initTransactionEnvolpeResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    initTransactionBodyResponse
}

type initTransactionBodyResponse struct {
	XMLName                    xml.Name `xml:"Body"`
	Fault                      *sign.SoapFault
	Ns2InitTransactionResponse *initTransactionResponse
}

type initTransactionResponse struct {
	XMLName xml.Name `xml:"initTransactionResponse"`
	Token   string   `xml:"return>token"`
	URL     string   `xml:"return>url"`
}

type transactionResultBodyRequest struct {
	XMLName                   xml.Name `xml:"soap:Body"`
	XMLnsSOAP                 string   `xml:"xmlns:soap,attr,omitempty"`
	ID                        string   `xml:"Id,attr,omitempty"`
	TnsAcknowledgeTransaction transactionResultResquest
}

type transactionResultResquest struct {
	XMLName    xml.Name `xml:"tns:getTransactionResult"`
	XMLnsTns   string   `xml:"xmlns:tns,attr,omitempty"`
	TokenInput string   `xml:"tokenInput"`
}

type transactionResultEnvolpeResponse struct {
	XMLName xml.Name                      `xml:"Envelope"`
	Body    transactionResultBodyResponse `xml:"Body"`
}

type transactionResultBodyResponse struct {
	XMLName                      xml.Name `xml:"Body"`
	Fault                        *sign.SoapFault
	Ns2TransactionResultResponse *transactionResultResponse
}

type transactionResultResponse struct {
	XMLName            xml.Name `xml:"getTransactionResultResponse"`
	AccountingDate     string   `xml:"return>accountingDate"`
	BuyOrder           string   `xml:"return>buyOrder"`
	CardNumber         string   `xml:"return>cardDetail>cardNumber"`
	CardExpirationDate string   `xml:"return>cardDetail>cardExpirationDate"`
	SharesNumber       int      `xml:"return>detailOutput>sharesNumber"`
	Amount             float64  `xml:"return>detailOutput>amount"`
	CommerceCode       string   `xml:"return>detailOutput>commerceCode"`
	DetailBuyOrder     string   `xml:"return>detailOutput>buyOrder"`
	AuthorizationCode  string   `xml:"return>detailOutput>authorizationCode"`
	PaymentTypeCode    string   `xml:"return>detailOutput>paymentTypeCode"`
	ResponseCode       string   `xml:"return>detailOutput>responseCode"`
	SessionID          string   `xml:"return>sessionId"`
	TransactionDate    string   `xml:"return>transactionDate"`
	URLRedirection     string   `xml:"return>urlRedirection"`
	VCI                string   `xml:"return>VCI"`
}
