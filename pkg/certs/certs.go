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

func NewCSR(CommonName string) (string,error) {

	buffer := new(bytes.Buffer)

	keyBytes, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return "", err
	}
	subj := pkix.Name{
		CommonName:         CommonName,
		Organization:       []string{"WATCHTOWER"},
		OrganizationalUnit: []string{"CD"},
		Locality:           []string{"Warwick"},
		Province:           []string{"New York"},
		Country:            []string{"US"},
	}

	rawSubj := subj.ToRDNSequence()

	asn1Subj, err := asn1.Marshal(rawSubj)

	if err != nil {
		return "", err
	}
	template := x509.CertificateRequest{
		RawSubject:         asn1Subj,
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, keyBytes)
	if err != nil {
		return "", err
	}
	pem.Encode(buffer, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	return buffer.String(), nil
}
