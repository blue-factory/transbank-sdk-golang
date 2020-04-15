package service

import (
	"encoding/xml"

	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

const (
	transactionType = "TR_NORMAL_WS"
)

// plusNormalBodyRequest ...
type plusNormalBodyRequest struct {
	XMLName            xml.Name `xml:"soap:Body"`
	XMLnsSOAP          string   `xml:"xmlns:soap,attr,omitempty"`
	ID                 string   `xml:"Id,attr,omitempty"`
	TnsInitTransaction plusNormalInitTransactionResquest
}

// plusNormalInitTransactionResquest ...
type plusNormalInitTransactionResquest struct {
	XMLName           xml.Name `xml:"tns:initTransaction"`
	XMLnsTns          string   `xml:"xmlns:tns,attr,omitempty"`
	SessionID         string   `xml:"wsInitTransactionInput>sessionId"`
	ReturnURL         string   `xml:"wsInitTransactionInput>returnURL"`
	FinalURL          string   `xml:"wsInitTransactionInput>finalURL"`
	BuyOrder          int64    `xml:"wsInitTransactionInput>buyOrder"`
	CommerceCode      int64    `xml:"wsInitTransactionInput>transactionDetails>commerceCode"`
	Amount            float64  `xml:"wsInitTransactionInput>transactionDetails>amount"`
	DetailBuyOrder    int64    `xml:"wsInitTransactionInput>transactionDetails>buyOrder"`
	WSTransactionType string   `xml:"wsInitTransactionInput>wSTransactionType"`
}

// plusNormalEnvolpeResponse ...
type plusNormalEnvolpeResponse struct {
	XMLName   xml.Name               `xml:"Envelope"`
	XMLnsSoap string                 `xml:"soap,attr"`
	Body      plusNormalBodyResponse `xml:"Body"`
}

// plusNormalBodyResponse ...
type plusNormalBodyResponse struct {
	XMLName                    xml.Name `xml:"Body"`
	XMLnsWsu                   string   `xml:"wsu,attr"`
	WsuID                      string   `xml:"Id,attr"`
	Ns2InitTransactionResponse plusNormalInitTransactionResponse
}

// plusNormalInitTransactionResponse ...
type plusNormalInitTransactionResponse struct {
	XMLName  xml.Name `xml:"initTransactionResponse"`
	XMLnsNs2 string   `xml:"ns2,attr"`
	Token    string   `xml:"return>token"`
	URL      string   `xml:"return>url"`
}

// ParamsPlusNormal ...
type ParamsPlusNormal struct {
	Amount    float64
	BuyOrder  int64
	SessionID string
	ReturnURL string
	FinalURL  string
}

// ResponsePlusNormal ...
type ResponsePlusNormal struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

// PlusNormal ...
type PlusNormal struct {
	Webpay *webpay.Webpay
}

// InitTransaction ...
func (pn *PlusNormal) InitTransaction(params ParamsPlusNormal) (*ResponsePlusNormal, error) {
	bodyRequest := plusNormalBodyRequest{
		ID:        "_0",
		XMLnsSOAP: "http://schemas.xmlsoap.org/soap/envelope/",
		TnsInitTransaction: plusNormalInitTransactionResquest{
			XMLnsTns:          "http://service.wswebpay.webpay.transbank.com/",
			SessionID:         params.SessionID,
			ReturnURL:         params.ReturnURL,
			FinalURL:          params.FinalURL,
			CommerceCode:      pn.Webpay.Config.CommerceCode,
			Amount:            params.Amount,
			BuyOrder:          params.BuyOrder,
			DetailBuyOrder:    params.BuyOrder,
			WSTransactionType: transactionType,
		},
	}

	b, err := pn.Webpay.SOAP(bodyRequest)
	if err != nil {
		return nil, err
	}

	res := &plusNormalEnvolpeResponse{}
	err = xml.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	return &ResponsePlusNormal{
		URL:   res.Body.Ns2InitTransactionResponse.URL,
		Token: res.Body.Ns2InitTransactionResponse.Token,
	}, nil
}

// GetTransactionResult ...
func (pn *PlusNormal) GetTransactionResult(token string) {}

/**************************************************/

// GetPlusNormal ...
func GetPlusNormal(w *webpay.Webpay) *PlusNormal {
	return &PlusNormal{
		Webpay: w,
	}
}
