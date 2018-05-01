package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"github.com/srleyva/CertAPI/pkg/certs"
)

func TestMyRouterAndHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/pki/NewCSR", strings.NewReader("google.com"))
	res := httptest.NewRecorder()
	NewPkiServer().ServeHTTP(res, req)

	if !strings.Contains(res.Body.String(), "CERTIFICATE REQUEST") {
		t.Errorf("cert not generated correctly: %s", res.Body.String())
	}
}

func TestNewCSRHandler(t *testing.T) {
	t.Run("Test Correct func is writing", func(t *testing.T) {
		certs.Initialize()
		req, _ := http.NewRequest("POST", "/pki/NewCSR", strings.NewReader("google.com"))
		res := httptest.NewRecorder()

		NewCSRHandler(res, req)

		if !strings.Contains(res.Body.String(), "CERTIFICATE REQUEST") {
			t.Errorf("Error: Invalid Config")
		}
	})
}

func TestGetConfigHandler(t *testing.T) {
	t.Run("Test Check config handler", func(t *testing.T) {
		certs.Initialize()
		req, _ := http.NewRequest("GET", "/pki/config", nil)
		res := httptest.NewRecorder()

		ConfigHandler(res, req)

		if !strings.Contains(res.Body.String(), "\"CA\":\"will.therealpki.com\"") {
			t.Errorf("Error: Response is malformed: %s", res.Body)
		}
	})
}