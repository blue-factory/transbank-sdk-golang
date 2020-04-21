package sign

type SOAPSigner interface {
	Sign(interface{}) ([]byte, error)
}

func New(privateKey, cert string) SOAPSigner {
	return &defaultSigner{key: privateKey, cert: cert}
}

type defaultSigner struct {
	key  string
	cert string
}

func (ds *defaultSigner) Sign(payload interface{}) ([]byte, error) {
	return ds.generateXMLRequest(payload)
}
