package configuration

import "errors"

// Configuration ...
type Configuration struct {
	PrivateCert   string
	PublicCert    string
	CommerceCode  int64
	CommerceEmail string
}

// New ...
func New(c Configuration) (*Configuration, error) {
	// validate private cert
	if c.PrivateCert == "" {
		return nil, errors.New("undefined Configuration.PrivateCert")
	}
	// TODO: check if a valid private cert
	if false {
		return nil, errors.New("invalid Configuration.PrivateCert")
	}

	// validate public cert
	if c.PrivateCert == "" {
		return nil, errors.New("undefined Configuration.PrivateCert")
	}
	// TODO: check if a valid public cert
	if false {
		return nil, errors.New("invalid Configuration.PrivateCert")
	}

	// validate commerce code
	if c.PrivateCert == "" {
		return nil, errors.New("undefined Configuration.CommerceCode")
	}

	// validate commerce email
	if c.PrivateCert == "" {
		return nil, errors.New("undefined Configuration.CommerceEmail")
	}

	return &Configuration{
		PrivateCert:   c.PrivateCert,
		PublicCert:    c.PublicCert,
		CommerceCode:  c.CommerceCode,
		CommerceEmail: c.CommerceEmail,
	}, nil
}
