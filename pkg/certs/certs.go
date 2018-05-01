// Handles the actual operations of dealing with certificates
package certs

import (
	"crypto/x509/pkix"
	"crypto/rsa"
	"crypto/rand"
	"encoding/asn1"
	"crypto/x509"
	"encoding/pem"
	"bytes"
)

// Main PKI config
var PKIConfig *Config

// Loads PKI config into application
func Initialize() error {
	var err error
	PKIConfig, err = NewConfig([]string{"../../", "./", "/etc/.pkiconf/", "$HOME/.pkiconf/"})
	return err
}

func GetConfig() *Config {
	return PKIConfig
}

// Generates a new Certificate Sign Request
func NewCSR(CommonName string) ([]byte,error) {

	buffer := new(bytes.Buffer)

	keyBytes, err := rsa.GenerateKey(rand.Reader, 2048)


	subj := pkix.Name{
		CommonName:         CommonName,
		Organization:       []string{PKIConfig.PKI.Organization},
		OrganizationalUnit: []string{PKIConfig.PKI.OrganizationalUnit},
		Locality:           []string{PKIConfig.PKI.Locality},
		Province:           []string{PKIConfig.PKI.Province},
		Country:            []string{PKIConfig.PKI.Country},
	}

	rawSubj := subj.ToRDNSequence()

	asn1Subj, err := asn1.Marshal(rawSubj)

	if err != nil {
		return nil, err
	}
	template := x509.CertificateRequest{
		RawSubject:         asn1Subj,
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, keyBytes)
	if err != nil {
		return nil, err
	}
	pem.Encode(buffer, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	return buffer.Bytes(), nil
}
