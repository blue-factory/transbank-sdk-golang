package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
)

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
