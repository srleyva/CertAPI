package certs

import (
	"testing"
)

func TestNewCSR(t *testing.T) {
	t.Run("test NewCSR", func(t *testing.T) {
		var csr string
		var err error
		if csr, err = NewCSR("jw.org"); err != nil || csr == "" {
			t.Error("CSR gen failed")
		}
	})
}
