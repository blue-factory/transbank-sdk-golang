package webpay

import "github.com/microapis/transbank-sdk-golang/pkg/sign"

// Webpay holds configuration that will be used by the services `service`
type Webpay struct {
	sign.SOAPSigner
	config *configuration
}

func (w *Webpay) GetCommerceCode() int64 {
	return w.config.CommerceCode
}

// New returns a configured Webpay instance
func New(privateCert, publicCert string, commerceCode int64, commerceEmail, service, environment string) (*Webpay, error) {
	c, err := newConfiguration(privateCert, publicCert, commerceCode, commerceEmail, service, environment)
	if err != nil {
		return nil, err
	}

	return new(c), nil
}

func new(c *configuration) *Webpay {
	w := &Webpay{
		config:     c,
		SOAPSigner: sign.New(c.PrivateCert, c.PublicCert),
	}

	return w
}
