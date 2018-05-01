package certs

import (
	"testing"
	"reflect"
)

func TestNewCSR(t *testing.T) {
	t.Run("test NewCSR", func(t *testing.T) {
		Initialize()
		var csr []byte
		var err error
		if csr, err = NewCSR("google.com"); err != nil || csr == nil {
			t.Errorf("CSR gen failed: %s", err)
		}
	})
}

func TestGetConfig(t *testing.T) {
	t.Run("test GetConfig", func(t *testing.T) {
		Initialize()
		actual := GetConfig()
		expected := PKIConfig

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("bad config returned: %s", actual)
		}
	})
}
