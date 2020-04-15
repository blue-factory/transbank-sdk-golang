package webpay

import (
	"errors"
	"fmt"

	"github.com/microapis/transbank-sdk-golang/pkg/configuration"
)

const (
	environmentIntegration   = "integration"
	environmentCertification = "certification"
	environmentTest          = "test"
	environmentLive          = "live"
	environmentProduction    = "production"

	serviceNormal       = "normal"
	serviceMallNormal   = "mallNormal"
	serviceCapture      = "capture"
	serviceNullify      = "nullify"
	serviceOneClick     = "oneClick"
	serviceOneClickMall = "oneClickMall"
)

// Webpay ...
type Webpay struct {
	Config configuration.Configuration

	service     string
	environment string
	wsdlURL     string
}

// New ...
func New(c configuration.Configuration) *Webpay {
	w := &Webpay{
		Config:      c,
		service:     serviceNormal,
		environment: environmentIntegration,
		wsdlURL:     buildWsdlURL(environmentIntegration, serviceNormal),
	}

	return w
}

// SetConfiguration ...
func (w *Webpay) SetConfiguration(c configuration.Configuration) {
	w.Config = c
}

// SetService ...
func (w *Webpay) SetService(service string) error {
	switch service {
	case serviceCapture, serviceMallNormal, serviceNormal, serviceNullify, serviceOneClick, serviceOneClickMall:
		w.service = service
	default:
		err := fmt.Sprintf("invalid service value: %v", service)
		return errors.New(err)
	}

	w.wsdlURL = buildWsdlURL(w.environment, w.service)

	return nil
}

// SetEnvironment ...
func (w *Webpay) SetEnvironment(environment string) error {
	switch environment {
	case environmentCertification, environmentIntegration, environmentLive, environmentProduction, environmentTest:
		w.environment = environment
	default:
		err := fmt.Sprintf("invalid environment value: %v", environment)
		return errors.New(err)
	}

	w.wsdlURL = buildWsdlURL(w.environment, w.service)

	return nil
}

// // GetTransbankCert ...
// func (w *Webpay) GetTransbankCert() string {
// 	tc := integration.TransbankCert

// 	switch w.environment {
// 	case environmentIntegration, environmentTest, environmentCertification:
// 		tc = integration.TransbankIntegrationCert
// 	}

// 	return tc
// }

// Clone ...
func (w *Webpay) Clone() Webpay {
	clone := Webpay{
		Config:      w.Config,
		environment: w.environment,
		service:     w.service,
		wsdlURL:     w.wsdlURL,
	}

	return clone
}
