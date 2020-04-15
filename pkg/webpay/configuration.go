package webpay

import (
	"errors"
	"fmt"
)

const (
	EnvironmentIntegration   = "integration"
	EnvironmentCertification = "certification"
	EnvironmentTest          = "test"
	EnvironmentLive          = "live"
	EnvironmentProduction    = "production"

	ServiceNormal       = "normal"
	ServiceMallNormal   = "mallNormal"
	ServiceCapture      = "capture"
	ServiceNullify      = "nullify"
	ServiceOneClick     = "oneClick"
	ServiceOneClickMall = "oneClickMall"
	ServicePatpass      = "patpass"
)

type configuration struct {
	PrivateCert   string
	PublicCert    string
	CommerceCode  int64
	CommerceEmail string
	Service       string
	Environment   string
}

func newConfiguration(privateCert, publicCert string, commerceCode int64, commerceEmail, service, environment string) (*configuration, error) {
	// validate private cert
	if privateCert == "" {
		return nil, errors.New("undefined Configuration.PrivateCert")
	}
	// TODO: check if a valid private cert
	if false {
		return nil, errors.New("invalid Configuration.PrivateCert")
	}

	// validate public cert
	if publicCert == "" {
		return nil, errors.New("undefined Configuration.PrivateCert")
	}
	// TODO: check if a valid public cert
	if false {
		return nil, errors.New("invalid Configuration.PrivateCert")
	}

	// validate commerce code
	if privateCert == "" {
		return nil, errors.New("undefined Configuration.CommerceCode")
	}

	// validate commerce email
	if privateCert == "" {
		return nil, errors.New("undefined Configuration.CommerceEmail")
	}

	// validate service
	switch service {
	case ServiceCapture, ServiceMallNormal, ServiceNormal, ServiceNullify, ServiceOneClick, ServiceOneClickMall:
	default:
		err := fmt.Sprintf("invalid service value: %v", service)
		return nil, errors.New(err)
	}

	// validate environment
	switch environment {
	case EnvironmentCertification, EnvironmentIntegration, EnvironmentLive, EnvironmentProduction, EnvironmentTest:
	default:
		err := fmt.Sprintf("invalid environment value: %v", environment)
		return nil, errors.New(err)
	}

	return &configuration{
		PrivateCert:   privateCert,
		PublicCert:    publicCert,
		CommerceCode:  commerceCode,
		CommerceEmail: commerceEmail,
		Service:       service,
		Environment:   environment,
	}, nil
}

func getIntegrationPlusNormal() *configuration {
	return &configuration{
		PrivateCert:   integrationPlusNormalPrivateCert,
		PublicCert:    integrationPlusNormalPublicCert,
		CommerceCode:  597020000540,
		CommerceEmail: "",
		Service:       ServiceNormal,
		Environment:   EnvironmentIntegration,
	}
}

func getIntegrationPlusMall() *configuration {
	return &configuration{
		PrivateCert:   integrationPlusMallPrivateCert,
		PublicCert:    integrationPlusMallPublicCert,
		CommerceCode:  597044444401,
		CommerceEmail: "",
		Service:       ServiceMallNormal,
		Environment:   EnvironmentIntegration,
	}
}

func getIntegrationPlusCapture() *configuration {
	return &configuration{
		PrivateCert:   integrationPlusCapturePrivateCert,
		PublicCert:    integrationPlusCapturePublicCert,
		CommerceCode:  597044444404,
		CommerceEmail: "",
		Service:       ServiceCapture,
		Environment:   EnvironmentIntegration,
	}
}

func getIntegrationOneClickNormal() *configuration {
	return &configuration{
		PrivateCert:   integrationOneClickNormalPrivateCert,
		PublicCert:    integrationOneClickNormalPublicCert,
		CommerceCode:  597044444405,
		CommerceEmail: "",
		Service:       ServiceOneClick,
		Environment:   EnvironmentIntegration,
	}
}

func getIntegrationPatpassNormal() *configuration {
	return &configuration{
		PrivateCert:   integrationPatpassNormalPrivateCert,
		PublicCert:    integrationPatpassNormalPublicCert,
		CommerceCode:  597020000548,
		CommerceEmail: "",
		Service:       ServicePatpass,
		Environment:   EnvironmentIntegration,
	}
}
