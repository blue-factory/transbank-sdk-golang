package webpay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"strings"
)

// envolpeRequest
type envolpeRequest struct {
	XMLName   xml.Name    `xml:"soap:Envelope"`
	XMLnsSoap string      `xml:"xmlns:soap,attr"`
	XMLnsTns  string      `xml:"xmlns:tns,attr"`
	XMLnsXsi  string      `xml:"xmlns:xsi,attr"`
	Header    headRequest `xml:"soap:Header"`
	Body      interface{} `xml:"soap:Body,omitempty"`
}

// headRequest ...
type headRequest struct {
	XMLName      xml.Name      `xml:"soap:Header"`
	WsseSecurity *wsseSecurity `xml:"wsse:Security,omitempty"`
}

// wsseSecurity ...
type wsseSecurity struct {
	XMLName            xml.Name   `xml:"wsse:Security"`
	XMLnsWsse          string     `xml:"xmlns:wsse,attr"`
	WsseMustUnderstand string     `xml:"wsse:mustUnderstand,attr"`
	X509Data           x509Data   `xml:"KeyInfo>X509Data"`
	Signature          *signature `xml:"Signature,omitempty"`
}

// x509Data ...
type x509Data struct {
	XMLName      xml.Name `xml:"X509Data"`
	XMLnsDs      string   `xml:"xmlns:ds,attr,omitempty"`
	IssuerName   string   `xml:"X509IssuerSerial>X509IssuerName"`
	SerialNumber string   `xml:"X509IssuerSerial>X509SerialNumber"`
	Certificate  string   `xml:"X509Certificate"`
}

// signature ...
type signature struct {
	XMLName        xml.Name   `xml:"Signature"`
	XMLns          string     `xml:"xmlns,attr"`
	SignedInfo     signedInfo `xml:"SignedInfo"`
	SignatureValue string     `xml:"SignatureValue"`
	X509Data       x509Data   `xml:"KeyInfo>wsse:SecurityTokenReference>X509Data,omitempty"`
}

// signedInfo ...
type signedInfo struct {
	XMLName                xml.Name               `xml:"SignedInfo"`
	XMLns                  string                 `xml:"xmlns,attr,omitempty"`
	CanonicalizationMethod canonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        signatureMethod        `xml:"SignatureMethod"`
	Reference              reference              `xml:"Reference,omitempty"`
}

// canonicalizationMethod ...
type canonicalizationMethod struct {
	XMLName   xml.Name `xml:"CanonicalizationMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

// signatureMethod ...
type signatureMethod struct {
	XMLName   xml.Name `xml:"SignatureMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

// reference ...
type reference struct {
	XMLName      xml.Name    `xml:"Reference"`
	URI          string      `xml:"URI,attr"`
	Transforms   []transform `xml:"Transforms>Transform"`
	DigestMethod digestMethod
	DigestValue  string `xml:"DigestValue"`
}

// transform ...
type transform struct {
	XMLName   xml.Name `xml:"Transform"`
	Algorithm string   `xml:"Algorithm,attr"`
}

// digestMethod ...
type digestMethod struct {
	XMLName   xml.Name `xml:"DigestMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

// GetXMLRequest ...
func (w *Webpay) GetXMLRequest(payload interface{}) ([]byte, error) {
	// decode and parse public cert
	block, _ := pem.Decode([]byte(w.Config.PublicCert))
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	// sanitilize certificate value
	public := strings.ReplaceAll(w.Config.PublicCert, "-----BEGIN CERTIFICATE-----", "")
	public = strings.ReplaceAll(public, "-----END CERTIFICATE-----", "")
	public = strings.ReplaceAll(public, "\r\n", "")
	public = strings.ReplaceAll(public, "\n", "")
	public = strings.ReplaceAll(public, "\r", "")

	digestValue, err := w.DigestValue(payload)
	if err != nil {
		return nil, err
	}

	signatureValue, err := w.SignatureValue(digestValue)
	if err != nil {
		return nil, err
	}

	// prepare <soap:Head>
	head := headRequest{
		WsseSecurity: &wsseSecurity{
			XMLnsWsse:          "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
			WsseMustUnderstand: "1",
			X509Data: x509Data{
				XMLnsDs:      "http://www.w3.org/2000/09/xmldsig#",
				IssuerName:   cert.Issuer.String(),
				SerialNumber: cert.SerialNumber.String(),
				Certificate:  public,
			},
			Signature: &signature{
				XMLns: "http://www.w3.org/2000/09/xmldsig#",
				SignedInfo: signedInfo{
					CanonicalizationMethod: canonicalizationMethod{
						Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
					},
					SignatureMethod: signatureMethod{
						Algorithm: "http://www.w3.org/2000/09/xmldsig#rsa-sha1",
					},
					Reference: reference{
						URI: "#_0",
						Transforms: []transform{
							transform{
								Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
							},
							transform{
								Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
							},
						},
						DigestMethod: digestMethod{
							Algorithm: "http://www.w3.org/2000/09/xmldsig#sha1",
						},
						DigestValue: digestValue,
					},
				},
				SignatureValue: signatureValue,
				X509Data: x509Data{
					XMLnsDs:      "http://www.w3.org/2000/09/xmldsig#",
					IssuerName:   cert.Issuer.String(),
					SerialNumber: cert.SerialNumber.String(),
					Certificate:  public,
				},
			},
		},
	}

	envolpe := envolpeRequest{
		XMLnsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XMLnsTns:  "http://service.wswebpay.webpay.transbank.com/",
		XMLnsXsi:  "http://www.w3.org/2001/XMLSchema-instance",
		Header:    head,
		Body:      payload,
	}

	parse, err := xml.Marshal(envolpe)
	if err != nil {
		return nil, err
	}

	out := xml.Header + string(parse)
	return []byte(out), nil
}

// DigestValue ...
func (w *Webpay) DigestValue(body interface{}) (string, error) {
	parse, err := xml.Marshal(body)
	if err != nil {
		return "", err
	}

	hash, err := hashSHA1(parse)
	if err != nil {
		return "", err
	}

	b64 := base64.StdEncoding.EncodeToString(hash)

	return b64, nil
}

// SignatureValue ...
func (w *Webpay) SignatureValue(digest string) (string, error) {
	s := signedInfo{
		XMLns: "http://www.w3.org/2000/09/xmldsig#",
		CanonicalizationMethod: canonicalizationMethod{
			Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
		},
		SignatureMethod: signatureMethod{
			Algorithm: "http://www.w3.org/2000/09/xmldsig#rsa-sha1",
		},
		Reference: reference{
			URI: "#_0",
			Transforms: []transform{
				transform{
					Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
				},
				transform{
					Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
				},
			},
			DigestMethod: digestMethod{
				Algorithm: "http://www.w3.org/2000/09/xmldsig#sha1",
			},
			DigestValue: digest,
		},
	}

	parse, err := xml.Marshal(s)
	if err != nil {
		return "", err
	}

	hash, err := hashRSASha1(parse, w.Config.PrivateCert)
	if err != nil {
		return "", err
	}

	b64 := base64.StdEncoding.EncodeToString(hash)

	return b64, nil
}

////////////////////////////////////////////////////////////////////////

func hashSHA1(value []byte) ([]byte, error) {
	hash := sha1.New()
	_, err := hash.Write(value)
	if err != nil {
		return nil, err
	}

	sum := hash.Sum(nil)

	return sum, nil
}

func hashRSASha1(value []byte, privateKey string) ([]byte, error) {
	hash, err := hashSHA1(value)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(privateKey))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hash)
	if err != nil {
		return nil, err
	}

	return sign, nil
}
