package service

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

const (
	transactionType = "TR_NORMAL_WS"
)

// ResponsePlusNormalInitTransaction represents the "initTransactionResponse" SOAP server response
type ResponsePlusNormalInitTransaction struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

// ResponsePlusNormalTransactionResult represents the "getTransactionResultResponse" SOAP server response
type ResponsePlusNormalTransactionResult struct {
	AccountingDate  string       `json:"accounting_date"`
	BuyOrder        string       `json:"buy_order"`
	CardDetail      cardDetail   `json:"card_detail"`
	DetailOutput    detailOutput `json:"detail_output"`
	SessionID       string       `json:"session_id"`
	TransactionDate string       `json:"transaction_date"`
	URLRedirection  string       `json:"url_redirection"`
	VCI             string       `json:"vci"`
}

type cardDetail struct {
	CardNumber string `json:"card_number"`
}

type detailOutput struct {
	SharesNumber      int     `json:"shares_number"`
	Amount            float64 `json:"amount"`
	CommerceCode      string  `json:"commerce_code"`
	BuyOrder          string  `json:"buy_order"`
	AuthorizationCode string  `json:"authorization_code"`
	PaymentTypeCode   string  `json:"payment_type_code"`
	ResponseCode      string  `json:"response_code"`
}

type plusNormal struct {
	webpay *webpay.Webpay
}

// GetPlusNormal returns an instance of Webpay's "plusNormal" service
func GetPlusNormal(w *webpay.Webpay) *plusNormal {
	return &plusNormal{
		webpay: w,
	}
}

// InitTransaction performans a "plusNormal" transaction and returns a token
func (pn *plusNormal) InitTransaction(amount float64, sessionID, buyOrder, returnURL, finalURL string) (*ResponsePlusNormalInitTransaction, error) {
	bodyRequest := plusNormalInitTransactionBodyRequest{
		ID:        "_0",
		XMLnsSOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		TnsInitTransaction: plusNormalInitTransactionResquest{
			XMLnsTns:          "http://service.wswebpay.webpay.transbank.com/",
			SessionID:         sessionID,
			ReturnURL:         returnURL,
			FinalURL:          finalURL,
			CommerceCode:      pn.webpay.GetCommerceCode(),
			Amount:            amount,
			BuyOrder:          buyOrder,
			DetailBuyOrder:    buyOrder,
			WSTransactionType: transactionType,
		},
	}

	b, err := pn.webpay.SOAP(bodyRequest)
	if err != nil {
		return nil, err
	}

	res := &plusNormalEnvolpeInitTransactionEnvolpeResponse{}
	err = xml.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	if res.Body.Fault != nil {
		errMsg := fmt.Sprintf("Error: code=%s message=%s", res.Body.Fault.Code, res.Body.Fault.Message)
		return nil, errors.New(errMsg)
	}

	it := res.Body.Ns2InitTransactionResponse

	return &ResponsePlusNormalInitTransaction{
		URL:   it.URL,
		Token: it.Token,
	}, nil
}

// GetTransactionResult validates a transaction given a token
func (pn *plusNormal) GetTransactionResult(token string) (*ResponsePlusNormalTransactionResult, error) {
	bodyRequest := plusNormalTransactionResultBodyRequest{
		ID:        "_0",
		XMLnsSOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		TnsAcknowledgeTransaction: plusNormalTransactionResultResquest{
			XMLnsTns:   "http://service.wswebpay.webpay.transbank.com/",
			TokenInput: token,
		},
	}

	b, err := pn.webpay.SOAP(bodyRequest)
	if err != nil {
		return nil, err
	}

	res := &plusNormalTransactionResultEnvolpeResponse{}
	err = xml.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	if res.Body.Fault != nil {
		errMsg := fmt.Sprintf("Error: code=%s message=%s", res.Body.Fault.Code, res.Body.Fault.Message)
		return nil, errors.New(errMsg)
	}

	tr := res.Body.Ns2TransactionResultResponse

	return &ResponsePlusNormalTransactionResult{
		AccountingDate: tr.AccountingDate,
		BuyOrder:       tr.BuyOrder,
		CardDetail: cardDetail{
			CardNumber: tr.CardNumber,
		},
		DetailOutput: detailOutput{
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

type plusNormalInitTransactionBodyRequest struct {
	XMLName            xml.Name `xml:"soap:Body"`
	XMLnsSOAP          string   `xml:"xmlns:soap,attr,omitempty"`
	ID                 string   `xml:"Id,attr,omitempty"`
	TnsInitTransaction plusNormalInitTransactionResquest
}

type plusNormalInitTransactionResquest struct {
	XMLName           xml.Name `xml:"tns:initTransaction"`
	XMLnsTns          string   `xml:"xmlns:tns,attr,omitempty"`
	SessionID         string   `xml:"wsInitTransactionInput>sessionId"`
	ReturnURL         string   `xml:"wsInitTransactionInput>returnURL"`
	FinalURL          string   `xml:"wsInitTransactionInput>finalURL"`
	BuyOrder          string   `xml:"wsInitTransactionInput>buyOrder"`
	CommerceCode      int64    `xml:"wsInitTransactionInput>transactionDetails>commerceCode"`
	Amount            float64  `xml:"wsInitTransactionInput>transactionDetails>amount"`
	DetailBuyOrder    string   `xml:"wsInitTransactionInput>transactionDetails>buyOrder"`
	WSTransactionType string   `xml:"wsInitTransactionInput>wSTransactionType"`
}

type plusNormalEnvolpeInitTransactionEnvolpeResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    plusNormalInitTransactionBodyResponse
}

type plusNormalInitTransactionBodyResponse struct {
	XMLName                    xml.Name `xml:"Body"`
	Fault                      *webpay.SoapFault
	Ns2InitTransactionResponse *plusNormalInitTransactionResponse
}

type plusNormalInitTransactionResponse struct {
	XMLName xml.Name `xml:"initTransactionResponse"`
	Token   string   `xml:"return>token"`
	URL     string   `xml:"return>url"`
}

type plusNormalTransactionResultBodyRequest struct {
	XMLName                   xml.Name `xml:"soap:Body"`
	XMLnsSOAP                 string   `xml:"xmlns:soap,attr,omitempty"`
	ID                        string   `xml:"Id,attr,omitempty"`
	TnsAcknowledgeTransaction plusNormalTransactionResultResquest
}

type plusNormalTransactionResultResquest struct {
	XMLName    xml.Name `xml:"tns:getTransactionResult"`
	XMLnsTns   string   `xml:"xmlns:tns,attr,omitempty"`
	TokenInput string   `xml:"tokenInput"`
}

type plusNormalTransactionResultEnvolpeResponse struct {
	XMLName xml.Name                                `xml:"Envelope"`
	Body    plusNormalTransactionResultBodyResponse `xml:"Body"`
}

type plusNormalTransactionResultBodyResponse struct {
	XMLName                      xml.Name `xml:"Body"`
	Fault                        *webpay.SoapFault
	Ns2TransactionResultResponse *plusNormalTransactionResultResponse
}

type plusNormalTransactionResultResponse struct {
	XMLName           xml.Name `xml:"getTransactionResultResponse"`
	AccountingDate    string   `xml:"return>accountingDate"`
	BuyOrder          string   `xml:"return>buyOrder"`
	CardNumber        string   `xml:"return>cardDetail>cardNumber"`
	SharesNumber      int      `xml:"return>detailOutput>sharesNumber"`
	Amount            float64  `xml:"return>detailOutput>amount"`
	CommerceCode      string   `xml:"return>detailOutput>commerceCode"`
	DetailBuyOrder    string   `xml:"return>detailOutput>buyOrder"`
	AuthorizationCode string   `xml:"return>detailOutput>authorizationCode"`
	PaymentTypeCode   string   `xml:"return>detailOutput>paymentTypeCode"`
	ResponseCode      string   `xml:"return>detailOutput>responseCode"`
	SessionID         string   `xml:"return>sessionId"`
	TransactionDate   string   `xml:"return>transactionDate"`
	URLRedirection    string   `xml:"return>urlRedirection"`
	VCI               string   `xml:"return>VCI"`
}
