package webpay

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	soapURL    = "https://webpay3g%s.transbank.cl/%s/%s?wsdl"
	soapAction = `\"\"`

	wsdlUrlsNormal       = "WSWebpayService"
	wsdlUrlsMallNormal   = "WSWebpayService"
	wsdlUrlsCapture      = "WSCommerceIntegrationService"
	wsdlUrlsNullify      = "WSCommerceIntegrationService"
	wsdlUrlsOneClick     = "OneClickPaymentService"
	wsdlUrlsOneClickMall = "WSOneClickMulticodeService"
)

// SOAP ...
func (w *Webpay) SOAP(payload interface{}) ([]byte, error) {
	XMLReq, err := w.generateXMLRequest(payload)
	if err != nil {
		return nil, err
	}

	url := w.wsdlURL

	client := http.Client{
		Timeout: time.Duration(50 * time.Second),
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(XMLReq))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/xml")
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return b, nil
}

// buildWsdlURL ...
func buildWsdlURL(environment string, service string) string {
	e := ""
	switch environment {
	case environmentIntegration, environmentTest, environmentCertification:
		e = "int"
	}

	s := "WSWebpayTransaction/cxf"
	if service == serviceOneClick {
		s = "webpayserver/wswebpay"
	}

	var wsdl string
	switch service {
	case serviceCapture:
		wsdl = wsdlUrlsCapture
	case serviceMallNormal:
		wsdl = wsdlUrlsMallNormal
	case serviceNormal:
		wsdl = wsdlUrlsNormal
	case serviceNullify:
		wsdl = wsdlUrlsNullify
	case serviceOneClick:
		wsdl = wsdlUrlsOneClick
	case serviceOneClickMall:
		wsdl = wsdlUrlsOneClickMall
	}

	URL := fmt.Sprintf(soapURL, e, s, wsdl)

	return URL
}
