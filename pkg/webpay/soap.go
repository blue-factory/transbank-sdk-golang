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
	wsdlUrlsPatpass      = "WSWebpayService"
	wsdlUrlsCapture      = "WSCommerceIntegrationService"
	wsdlUrlsNullify      = "WSCommerceIntegrationService"
	wsdlUrlsOneClick     = "OneClickPaymentService"
	wsdlUrlsOneClickMall = "WSOneClickMulticodeService"
)

// SOAP This method performs a SOAP request to the server with a given payload
func (w *Webpay) SOAP(payload interface{}) ([]byte, error) {
	XMLReq, err := w.generateXMLRequest(payload)
	if err != nil {
		return nil, err
	}

	url := buildWsdlURL(w.config)

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

func buildWsdlURL(c *configuration) string {
	environment := c.Environment
	service := c.Service

	e := ""
	switch environment {
	case EnvironmentIntegration, EnvironmentTest, EnvironmentCertification:
		e = "int"
	}

	s := "WSWebpayTransaction/cxf"
	if service == ServiceOneClick {
		s = "webpayserver/wswebpay"
	}

	var wsdl string
	switch service {
	case ServiceCapture:
		wsdl = wsdlUrlsCapture
	case ServiceMallNormal:
		wsdl = wsdlUrlsMallNormal
	case ServicePatpass:
		wsdl = wsdlUrlsPatpass
	case ServiceNormal:
		wsdl = wsdlUrlsNormal
	case ServiceNullify:
		wsdl = wsdlUrlsNullify
	case ServiceOneClick:
		wsdl = wsdlUrlsOneClick
	case ServiceOneClickMall:
		wsdl = wsdlUrlsOneClickMall
	}

	URL := fmt.Sprintf(soapURL, e, s, wsdl)

	return URL
}
