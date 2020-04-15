package webpay

// Webpay ...
type Webpay struct {
	config  *configuration
	wsdlURL string
}

func (w *Webpay) GetPrivateCert() string {
	return w.config.PrivateCert
}

func (w *Webpay) GetPublicCert() string {
	return w.config.PublicCert
}

func (w *Webpay) GetCommerceCode() int64 {
	return w.config.CommerceCode
}

// New ...
func New(privateCert, publicCert string, commerceCode int64, commerceEmail, service, environment string) (*Webpay, error) {
	c, err := newConfiguration(privateCert, publicCert, commerceCode, commerceEmail, service, environment)
	if err != nil {
		return nil, err
	}

	return new(c), nil
}

// NewIntegrationPlusNormal ...
func NewIntegrationPlusNormal() *Webpay {
	return new(GetIntegrationPlusNormal())
}

func new(c *configuration) *Webpay {
	w := &Webpay{
		config:  c,
		wsdlURL: buildWsdlURL(c.Environment, c.Service),
	}

	return w
}
