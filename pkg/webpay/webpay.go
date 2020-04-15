package webpay

// Webpay holds configuration that will be used by the services `service`
type Webpay struct {
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

// NewIntegrationPlusNormal returns a configured Webpay instance that will use
// the integration environment
func NewIntegrationPlusNormal() *Webpay {
	return new(getIntegrationPlusNormal())
}

func new(c *configuration) *Webpay {
	w := &Webpay{
		config: c,
	}

	return w
}
